package shim

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric/core/chaincode/shim/crypto/attr"
)

type CustomMockStub struct {
	stub           *MockStub
	CertAttributes map[string][]byte
}

func NewCustomMockStub(name string, cc Chaincode, attributes map[string][]byte) *CustomMockStub {
	s := new(CustomMockStub)
	s.stub = NewMockStub(name, cc)
	s.CertAttributes = attributes
	return s
}

func (mock *CustomMockStub) ReadCertAttribute(attributeName string) ([]byte, error) {
	return mock.CertAttributes[attributeName], nil
}

func (mock *CustomMockStub) GetState(key string) ([]byte, error) {
	return mock.stub.GetState(key)
}

func (mock *CustomMockStub) GetTxID() string {
	return mock.stub.GetTxID()
}

func (mock *CustomMockStub) MockInit(uuid string, function string, args []string) ([]byte, error) {
	mock.stub.args = getBytes(function, args)
	mock.MockTransactionStart(uuid)
	bytes, err := mock.stub.cc.Init(mock, function, args)
	mock.MockTransactionEnd(uuid)
	return bytes, err
}

func (mock *CustomMockStub) MockInvoke(uuid string, function string, args []string) ([]byte, error) {
	mock.stub.args = getBytes(function, args)
	mock.MockTransactionStart(uuid)
	bytes, err := mock.stub.cc.Invoke(mock, function, args)
	mock.MockTransactionEnd(uuid)
	return bytes, err
}

func (mock *CustomMockStub) MockQuery(function string, args []string) ([]byte, error) {
	mock.stub.args = getBytes(function, args)
	bytes, err := mock.stub.cc.Query(mock, function, args)
	return bytes, err
}

func (mock *CustomMockStub) PutState(key string, value []byte) error {
	return mock.stub.PutState(key, value)
}

func (mock *CustomMockStub) MockTransactionStart(txid string) {
	mock.stub.MockTransactionStart(txid)
}

func (mock *CustomMockStub) MockTransactionEnd(uuid string) {
	mock.stub.MockTransactionEnd(uuid)
}

func (mock *CustomMockStub) GetArgs() [][]byte {
	return mock.stub.GetArgs()
}

func (mock *CustomMockStub) GetStringArgs() []string {
	return mock.stub.GetStringArgs()
}

func (mock *CustomMockStub) MockPeerChaincode(invokableChaincodeName string, otherStub *MockStub) {
	mock.stub.MockPeerChaincode(invokableChaincodeName, otherStub)
}

func (mock *CustomMockStub) DelState(key string) error {
	return mock.stub.DelState(key)
}

func (mock *CustomMockStub) RangeQueryState(startKey, endKey string) (StateRangeQueryIteratorInterface, error) {
	return mock.stub.RangeQueryState(startKey, endKey)
}

func (mock *CustomMockStub) CreateTable(name string, columnDefinitions []*ColumnDefinition) error {
	return nil
}

func (mock *CustomMockStub) GetTable(tableName string) (*Table, error) {
	return nil, nil
}

func (mock *CustomMockStub) DeleteTable(tableName string) error {
	return nil
}

func (mock *CustomMockStub) InsertRow(tableName string, row Row) (bool, error) {
	return false, nil
}

func (mock *CustomMockStub) ReplaceRow(tableName string, row Row) (bool, error) {
	return false, nil
}

func (mock *CustomMockStub) GetRow(tableName string, key []Column) (Row, error) {
	return mock.stub.GetRow(tableName, key)
}

func (mock *CustomMockStub) GetRows(tableName string, key []Column) (<-chan Row, error) {
	return nil, nil
}

func (mock *CustomMockStub) DeleteRow(tableName string, key []Column) error {
	return nil
}
func (mock *CustomMockStub) InvokeChaincode(chaincodeName string, args [][]byte) ([]byte, error) {
	return mock.stub.InvokeChaincode(chaincodeName, args)
}

func (mock *CustomMockStub) QueryChaincode(chaincodeName string, args [][]byte) ([]byte, error) {
	return mock.stub.QueryChaincode(chaincodeName, args)
}

func (mock *CustomMockStub) VerifyAttribute(attributeName string, attributeValue []byte) (bool, error) {
	return false, nil
}

func (mock *CustomMockStub) VerifyAttributes(attrs ...*attr.Attribute) (bool, error) {
	return false, nil
}

func (stub *CustomMockStub) VerifySignature(certificate, signature, message []byte) (bool, error) {
	return false, nil
}

func (mock *CustomMockStub) GetCallerCertificate() ([]byte, error) {
	return nil, nil
}

func (mock *CustomMockStub) GetCallerMetadata() ([]byte, error) {
	return nil, nil
}

func (mock *CustomMockStub) GetBinding() ([]byte, error) {
	return nil, nil
}

func (mock *CustomMockStub) GetPayload() ([]byte, error) {
	return nil, nil
}

func (mock *CustomMockStub) GetTxTimestamp() (*timestamp.Timestamp, error) {
	return nil, nil
}

func (mock *CustomMockStub) SetEvent(name string, payload []byte) error {
	return nil
}
