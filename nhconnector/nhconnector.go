package nhconnector

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/0xPolygonHermez/zkevm-node/aggregator/prover"
	"github.com/0xPolygonHermez/zkevm-node/log"
	gsrpc "github.com/centrifuge/go-substrate-rpc-client/v4"
	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/ybbus/jsonrpc/v3"
)

type NHConnector struct {
	nhApi     *gsrpc.SubstrateAPI
	nhRpc     jsonrpc.RPCClient
	nhKeyPair signature.KeyringPair
}

func New() (NHConnector, error) {

	// set websocket endoint
	nhWsURL := os.Getenv("NH_WS_URL")

	// Instantiate the API
	api, err := gsrpc.NewSubstrateAPI(nhWsURL)
	if err != nil {
		panic(err)
	}

	// Instantiate the RPC client
	nhRpcURL := os.Getenv("NH_RPC_URL")
	rpcClient := jsonrpc.NewClient(nhRpcURL)

	// Create key pair from seed
	keyPair, err := signature.KeyringPairFromSecret(os.Getenv("NH_SEED_PHRASE"), 42)

	nhConnector := NHConnector{
		nhApi:     api,
		nhRpc:     rpcClient,
		nhKeyPair: keyPair,
	}

	return nhConnector, nil
}

// Get the proof merkle path from the attestation from NH
func (nhConnect *NHConnector) GetProofMerklePath(attestationID types.U64, proof string) PoeProofPathRequestResponse {
	log.Debug("Get Proof merkle path  for proof: ", proof, " with attestation ID: ", attestationID)

	// Perform RPC Call
	response, err := nhConnect.nhRpc.Call(context.Background(), PoeProofPathRequest,
		attestationID,
		proof)
	if err != nil {
		log.Error(err)
	}
	log.Debug("PoeProofPath: ", response)

	if response.Error != nil {
		switch response.Error.Code {
		case 1:
			log.Error("ERROR in PoeProofPath: ", ErrProofNotFoundInAttestation)
		case 2:
			log.Error("ERROR in PoeProofPath: ", ErrAttestationNotPublishedYet)
		default:
			log.Error("ERROR in PoeProofPath: ", response.Error)
		}
	}

	var jsonResponse *PoeProofPathRequestResponse
	err = response.GetObject(&jsonResponse)
	if err != nil || jsonResponse == nil {
		// some error on json unmarshal level or json result field was null
		log.Error(err)
	}

	return *jsonResponse
}

// Send the final proof to NH
func (nhConnector *NHConnector) SendProofToNH(finalProof *prover.FinalProof) (PoeNewElement, error) {

	// Get latest metadata
	meta, err := nhConnector.nhApi.RPC.State.GetMetadataLatest()
	if err != nil {
		return PoeNewElement{}, err
	}

	proofString := finalProof.GetProof() + strings.ReplaceAll(finalProof.GetCompactPublicInput(), "0x", "")
	log.Info("Sending proof to L0...")
	// sign transaction
	log.Info("Sending proof...", proofString, "PI: ", strings.ReplaceAll(finalProof.GetCompactPublicInput(), "0x", ""))
	//Mocked Proof (ForkID 6)
	//proof, err := Bytes800FromHex("0x2ecc31435ec6d6963d463c38ea5662d9c94a67e441e7bc611598ebcc59f571880768291fd5d95fcf02bce7e4fde1f048b843bbffab1f242904e82d443a4ebb61150c3a4afdbb62d034320da390e3585a30ba13f4df73798b78e5a75655d3350d19fca02cc5838405f9ae4177ac7117971af2cb5006d7a46436f644410d6e7c52099f803c0f18d4b44fe22f3100d1fc80ccb7309fa7168f51bc64f3fc0f1dbd240b0573f3593238e56b23e75246d9d0f6f6a5cf824700667e3482ca9fecf74cdc0b308e6a8f69dccb9ca00d540543441f7030928da766406a152427bdf31b44051b6b5198b34006f9ac34c6c857e450cc11f5c6b6d21119fe283738581c0ad8bd0f8cbdbc574f64da884f6a02e00669f3eef10138266f3d7fa278aef1b1c60171005c0c2b8b2429c5003c5ab24af44cb1ab81cdc96dcaf6004a0f74406bb10f45233b13015cef8c40c491a7770efd0a8d8a64186d4f3827e74972bfc25b11f1f002550a5e253c923c5783026c7439601595477f1a212de449c64a8ae5e2fc0313127bf9cd5146217e531196ce65ccef3249375450d6932151f923c39e6a73588223d90f5bf230eee5a6cb6463f161602cc37fe538e2954ebef695b926b76e3fae299c60c1952aa4b1f246204ac7c22c0156ede30aeb73444ee40d69c0f131fa471fdca090abfd38541c88ee73624657a695155748643f7834b80d1c0481079e670033817252b24575a4e6007f08f37c34462d5e9fd50b1e83ec8cfc86149400d519e224ee11831ac393e3a09730be6f385ae5c9e14446fde5069fea751fb6b48211c85268b8017de7981eb1bd78526bc20d5f863ad3abe249728ca7b75b2146c1254465d6100a911213d95f800779e74f6701b1dfa0b6660642108fd2c7cd2f131d9163eeebe9d8aabdf8d37fde4451f762be478d117688e0a6ed2648dbe025e82a4b13ee629a73d1efa6f269747506058746aa589bb961c1385bb2b30e0086f010ef87535f2137a04f19fe5aa7c4f348c32ce6f5b0b45bb503895673a8a51d7f1b0228693fbfb38be718b04c9fdf116a97d7f30e670db84d21bb0d12fc57645415950a3fab52ee1557ac7b895deeca2eb27bacfc3b9e26a39b1875149680611d")
	proof, err := Bytes800FromHex(proofString)
	if err != nil {
		return PoeNewElement{}, err
	}

	c, err := types.NewCall(meta, SubmitProofExtrinsic, proof)

	if err != nil {
		return PoeNewElement{}, err
	}

	// Create the extrinsic
	ext := types.NewExtrinsic(c)

	genesisHash, err := nhConnector.nhApi.RPC.Chain.GetBlockHash(0)
	if err != nil {
		return PoeNewElement{}, err
	}

	rv, err := nhConnector.nhApi.RPC.State.GetRuntimeVersionLatest()
	if err != nil {
		return PoeNewElement{}, err
	}

	key, err := types.CreateStorageKey(meta, "System", "Account", nhConnector.nhKeyPair.PublicKey)
	if err != nil {
		return PoeNewElement{}, err
	}

	var accountInfo types.AccountInfo
	ok, err := nhConnector.nhApi.RPC.State.GetStorageLatest(key, &accountInfo)
	if err != nil || !ok {
		return PoeNewElement{}, err
	}

	nonce := uint32(accountInfo.Nonce)

	o := types.SignatureOptions{
		BlockHash:          genesisHash,
		Era:                types.ExtrinsicEra{IsMortalEra: false},
		GenesisHash:        genesisHash,
		Nonce:              types.NewUCompactFromUInt(uint64(nonce)),
		SpecVersion:        rv.SpecVersion,
		Tip:                types.NewUCompactFromUInt(0),
		TransactionVersion: rv.TransactionVersion,
	}

	log.Debug("Sending proof with to NH mainchain... ")

	// Sign the transaction
	err = ext.Sign(nhConnector.nhKeyPair, o)
	if err != nil {
		return PoeNewElement{}, err
	}

	// Do the transfer and track the actual status
	subEx, err := nhConnector.nhApi.RPC.Author.SubmitAndWatchExtrinsic(ext)
	var blockHash types.Hash

	if err != nil {
		return PoeNewElement{}, err
	}
	defer subEx.Unsubscribe()

	for {
		status := <-subEx.Chan()

		if status.IsInBlock {
			blockHash = status.AsInBlock
			break
		}
	}
	log.Debug("Transaction included in block: ", blockHash.Hex())
	block, err := nhConnector.nhApi.RPC.Chain.GetBlock(blockHash)
	if err != nil {
		return PoeNewElement{}, err
	}
	var extrinsicId uint32
	for pos, e := range block.Block.Extrinsics {
		if e.Signature.Signature.AsSr25519 == ext.Signature.Signature.AsSr25519 {
			extrinsicId = uint32(pos)
		}
	}

	// Subscribe to system events via storage
	keyEvents, err := types.CreateStorageKey(meta, "System", "Events", nil)
	if err != nil {
		return PoeNewElement{}, err
	}

	stdEvents := types.EventRecords{}
	hlEvents := NHEventsRecords{}
	records, err := nhConnector.nhApi.RPC.State.GetStorageRaw(keyEvents, blockHash)
	if err != nil {
		return PoeNewElement{}, err
	}

	err = DecodeEventRecords(types.EventRecordsRaw(*records), meta, &stdEvents, &hlEvents)
	if err != nil {
		return PoeNewElement{}, err
	}

	// Search the Event
	var proofVerifiedElement PoeNewElement
	for _, e := range hlEvents.Poe_NewElement {
		if e.Phase.IsApplyExtrinsic && e.Phase.AsApplyExtrinsic == extrinsicId {
			log.Debug("tSettelmentFFlonkPallet:ProofVerified, ProofValue: ", e.Value.Hex(), " AttestationID: ", e.AttestationId)
			proofVerifiedElement = e
		}
	}
	return proofVerifiedElement, err
}

func ensureValidStruct(tval reflect.Value) (reflect.Value, error) {
	none := reflect.Value{}
	// ensure t is a pointer
	ttyp := tval.Type()
	if ttyp.Kind() != reflect.Ptr {
		return none, fmt.Errorf("target must be a pointer, but is  %v ", ttyp)
	}
	// ensure t is not a nil pointer
	if tval.IsNil() {
		return none, fmt.Errorf("target is a nil pointer")
	}
	val := tval.Elem()
	typ := val.Type()
	// ensure val can be set
	if !val.CanSet() {
		return none, fmt.Errorf("unsettable value %v", typ)
	}
	// ensure val points to a array
	if val.Kind() != reflect.Struct {
		return none, fmt.Errorf("target must point to a struct, but is %v", typ)
	}
	return val, nil
}

func (s *Discover) FieldByName(name string) reflect.Value {
	for _, obj := range s.objs {
		if !obj.IsValid() {
			panic("No valid object")
		}
		val := obj.FieldByName(name)
		if val.IsValid() {
			return val
		}
	}
	return reflect.Value{}
}

func getDiscover(tStd interface{}, tExt interface{}) (*Discover, error) {
	objs := make([]reflect.Value, 2)
	val, err := ensureValidStruct(reflect.ValueOf(tStd))
	if err != nil {
		return nil, fmt.Errorf("[Interface #%v] %s", 0, err.Error())
	}
	if !val.IsValid() {
		return nil, fmt.Errorf("not valid value val")
	}
	objs[0] = val
	exVal, err := ensureValidStruct(reflect.ValueOf(tExt))
	if err != nil {
		return nil, fmt.Errorf("[Interface #%v] %s", 1, err.Error())
	}
	if !exVal.IsValid() {
		return nil, fmt.Errorf("not valid value exVal")
	}
	objs[1] = exVal
	return &Discover{objs: objs}, nil
}

// DecodeEventRecords decodes the events records from an EventRecordRaw into a target t using the given Metadata m
// If this method returns an error like `unable to decode Phase for event #x: EOF`, it is likely that you have defined
// a custom event record with a wrong type. For example your custom event record has a field with a length prefixed
// type, such as types.Bytes, where your event in reallity contains a fixed width type, such as a types.U32.
func DecodeEventRecords(e types.EventRecordsRaw, m *types.Metadata, tStd interface{}, tExt interface{}) error { //nolint:funlen

	discover, err := getDiscover(tStd, tExt)
	if err != nil {
		return err
	}

	decoder := scale.NewDecoder(bytes.NewReader(e))

	// determine number of events
	n, err := decoder.DecodeUintCompact()
	if err != nil {
		return err
	}

	// iterate over events
	for i := uint64(0); i < n.Uint64(); i++ {

		// decode Phase
		phase := types.Phase{}
		err := decoder.Decode(&phase)
		if err != nil {
			return fmt.Errorf("unable to decode Phase for event #%v: %v", i, err)
		} else {
			fmt.Printf("Phase %v\n", phase)
		}

		// decode EventID
		id := types.EventID{}
		err = decoder.Decode(&id)
		if err != nil {
			return fmt.Errorf("unable to decode EventID for event #%v: %v", i, err)
		}

		// ask metadata for method & event name for event
		moduleName, eventName, err := m.FindEventNamesForEventID(id)
		if err != nil {
			return fmt.Errorf("unable to find event with EventID %v in metadata for event #%v: %s", id, i, err)
		}

		// check whether name for eventID exists in t
		field := discover.FieldByName(fmt.Sprintf("%v_%v", moduleName, eventName))
		if !field.IsValid() {
			return fmt.Errorf("unable to find field %v_%v for event #%v with EventID %v", moduleName, eventName, i, id)
		} else {

			// create a pointer to with the correct type that will hold the decoded event
			holder := reflect.New(field.Type().Elem())

			// ensure first field is for Phase, last field is for Topics
			numFields := holder.Elem().NumField()
			if numFields < 2 {
				return fmt.Errorf("expected event #%v with EventID %v, field %v_%v to have at least 2 fields "+
					"(for Phase and Topics), but has %v fields", i, id, moduleName, eventName, numFields)
			}
			phaseField := holder.Elem().FieldByIndex([]int{0})
			if phaseField.Type() != reflect.TypeOf(phase) {
				return fmt.Errorf("expected the first field of event #%v with EventID %v, field %v_%v to be of type "+
					"types.Phase, but got %v", i, id, moduleName, eventName, phaseField.Type())
			}
			topicsField := holder.Elem().FieldByIndex([]int{numFields - 1})
			if topicsField.Type() != reflect.TypeOf([]types.Hash{}) {
				return fmt.Errorf("expected the last field of event #%v with EventID %v, field %v_%v to be of type "+
					"[]types.Hash for Topics, but got %v", i, id, moduleName, eventName, topicsField.Type())
			}

			// set the phase we decoded earlier
			phaseField.Set(reflect.ValueOf(phase))

			// set the remaining fields
			for j := 1; j < numFields; j++ {
				err = decoder.Decode(holder.Elem().FieldByIndex([]int{j}).Addr().Interface())
				if err != nil {
					return fmt.Errorf("unable to decode field %v event #%v with EventID %v, field %v_%v: %v", j, i, id, moduleName,
						eventName, err)
				}
			}

			// add the decoded event to the slice
			field.Set(reflect.Append(field, holder.Elem()))
		}
	}

	return nil
}

type Discover struct {
	objs []reflect.Value
}
