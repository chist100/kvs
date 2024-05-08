package keeper_test

import (
	"fmt"
	keepertest "kvs/testutil/keeper"
	"kvs/x/kvs/keeper"
	"kvs/x/kvs/types"
	"math/rand"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type (
	KvsTestSuite struct {
		suite.Suite
		goCtx     sdk.Context
		msgServer types.MsgServer
		kvsKeeper *keeper.Keeper
	}
)

const (
	testCreatorAddr     = "kvs1r38k70lalz8dahvg965rp7xw6v0g8asch6dlvv"
	testAdminFirstAddr  = "kvs1h8pqxellgwywtfkrfct6rr44zyuhhghxpyrsk9"
	testAdminSecondAddr = "kvs1fvgqul0m6tjkgndz2q292an5hhnhdvtem8777u"
	testAdminThirdAddr  = "kvs1mvqzw2h2kntqat4gdqlcgl4xe527qhgwhqfwgq"
)

func (suite *KvsTestSuite) SetupSuite() {
	k, goCtx := keepertest.KvsKeeper(suite.T())
	suite.goCtx = goCtx
	suite.kvsKeeper = k
	suite.msgServer = keeper.NewMsgServerImpl(*k)
}

func (suite *KvsTestSuite) SetupTest() {
}
func (suite *KvsTestSuite) TearDownTest() {
	suite.kvsKeeper.RemoveAcl(suite.goCtx)
	for _, data := range suite.kvsKeeper.GetAllData(suite.goCtx) {
		suite.kvsKeeper.RemoveData(suite.goCtx, data.Index)
	}
	for _, prop := range suite.kvsKeeper.GetAllProposal(suite.goCtx) {
		suite.kvsKeeper.RemoveData(suite.goCtx, prop.Index)
	}
}

func TestRunAll(t *testing.T) {
	suite.Run(t, new(KvsTestSuite))
}

func (suite *KvsTestSuite) TestKvs_AddressRegistration_Success() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testAddresses := []string{testAdminFirstAddr, testAdminSecondAddr, testAdminThirdAddr}
	msg := &types.MsgAddressRegistration{
		Creator:   testCreatorAddr,
		Addresses: testAddresses,
	}
	_, err := suite.msgServer.AddressRegistration(ctx, msg)
	if err != nil {
		suite.T().Fail()
	}

	acl, exuist := suite.kvsKeeper.GetAcl(suite.goCtx)
	assert.Equal(suite.T(), true, exuist)
	assert.Equal(suite.T(), testAddresses, acl.Adresses)
}

func (suite *KvsTestSuite) TestKvs_AddressRegistration_Fail() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testAddresses := []string{testAdminFirstAddr, testAdminSecondAddr}
	msg := &types.MsgAddressRegistration{
		Creator:   testCreatorAddr,
		Addresses: testAddresses,
	}
	_, err := suite.msgServer.AddressRegistration(ctx, msg)
	if err == nil {
		suite.T().Fail()
	}
}

func (suite *KvsTestSuite) TestKvs_AddressRegistration_AlreadySet_Fail() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testAddresses := []string{testAdminFirstAddr, testAdminSecondAddr, testAdminThirdAddr}
	msg := &types.MsgAddressRegistration{
		Creator:   testCreatorAddr,
		Addresses: testAddresses,
	}
	suite.kvsKeeper.SetAcl(suite.goCtx, types.Acl{
		Adresses: testAddresses,
	})
	_, err := suite.msgServer.AddressRegistration(ctx, msg)
	if err == nil {
		suite.T().Fail()
	}
}

func (suite *KvsTestSuite) TestKvs_DataProposal_Success() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testKey := fmt.Sprint("testKey%i", rand.Int())
	testValue := fmt.Sprint("testKey%i", rand.Int())
	msg := &types.MsgDataProposal{
		Creator: testCreatorAddr,
		Key:     testKey,
		Value:   testValue,
	}
	_, err := suite.msgServer.DataProposal(ctx, msg)
	if err != nil {
		suite.T().Fail()
	}
	var testAck []string
	prop, exuist := suite.kvsKeeper.GetProposal(suite.goCtx, testKey)
	assert.Equal(suite.T(), true, exuist)
	assert.Equal(suite.T(), testKey, prop.Index)
	assert.Equal(suite.T(), testValue, prop.Value)
	assert.Equal(suite.T(), testAck, prop.Acknowledgments)
}

func (suite *KvsTestSuite) TestKvs_DataProposal_ProposalAlreadyExist_Fail() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testKey := fmt.Sprint("testKey%i", rand.Int())
	testValue := fmt.Sprint("testKey%i", rand.Int())
	msg := &types.MsgDataProposal{
		Creator: testCreatorAddr,
		Key:     testKey,
		Value:   testValue,
	}
	suite.kvsKeeper.SetProposal(suite.goCtx, types.Proposal{
		Index:           testKey,
		Value:           testValue,
		Acknowledgments: []string{},
	})

	_, err := suite.msgServer.DataProposal(ctx, msg)
	if err == nil {
		suite.T().Fail()
	}
}

func (suite *KvsTestSuite) TestKvs_DataProposal_DataAlreadyExist_Fail() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testKey := fmt.Sprint("testKey%i", rand.Int())
	testValue := fmt.Sprint("testKey%i", rand.Int())
	msg := &types.MsgDataProposal{
		Creator: testCreatorAddr,
		Key:     testKey,
		Value:   testValue,
	}
	suite.kvsKeeper.SetData(suite.goCtx, types.Data{
		Index: testKey,
		Value: testValue,
	})

	_, err := suite.msgServer.DataProposal(ctx, msg)
	if err == nil {
		suite.T().Fail()
	}
}

func (suite *KvsTestSuite) TestKvs_DataConfirmation_Success() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testKey := fmt.Sprint("testKey%i", rand.Int())
	testValue := fmt.Sprint("testKey%i", rand.Int())

	testAddresses := []string{testAdminFirstAddr, testAdminSecondAddr, testAdminThirdAddr}
	suite.kvsKeeper.SetAcl(suite.goCtx, types.Acl{
		Adresses: testAddresses,
	})

	suite.kvsKeeper.SetProposal(suite.goCtx, types.Proposal{
		Index:           testKey,
		Value:           testValue,
		Acknowledgments: []string{},
	})

	msg := &types.MsgDataConfirmation{
		Creator: testAdminFirstAddr,
		Key:     testKey,
	}
	_, err := suite.msgServer.DataConfirmation(ctx, msg)
	if err != nil {
		suite.T().Fail()
	}
	testAck := []string{testAdminFirstAddr}
	prop, exuist := suite.kvsKeeper.GetProposal(suite.goCtx, testKey)
	assert.Equal(suite.T(), true, exuist)
	assert.Equal(suite.T(), testKey, prop.Index)
	assert.Equal(suite.T(), testValue, prop.Value)
	assert.Equal(suite.T(), testAck, prop.Acknowledgments)
}

func (suite *KvsTestSuite) TestKvs_DataConfirmation_AlreadyConfirmed_Fail() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testKey := fmt.Sprint("testKey%i", rand.Int())
	testValue := fmt.Sprint("testKey%i", rand.Int())

	testAddresses := []string{testAdminFirstAddr, testAdminSecondAddr, testAdminThirdAddr}
	suite.kvsKeeper.SetAcl(suite.goCtx, types.Acl{
		Adresses: testAddresses,
	})

	suite.kvsKeeper.SetProposal(suite.goCtx, types.Proposal{
		Index:           testKey,
		Value:           testValue,
		Acknowledgments: []string{testAdminFirstAddr},
	})

	msg := &types.MsgDataConfirmation{
		Creator: testAdminFirstAddr,
		Key:     testKey,
	}
	_, err := suite.msgServer.DataConfirmation(ctx, msg)
	if err == nil {
		suite.T().Fail()
	}
}

func (suite *KvsTestSuite) TestKvs_DataConfirmation_NotExist_Fail() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testKey := fmt.Sprint("testKey%i", rand.Int())

	testAddresses := []string{testAdminFirstAddr, testAdminSecondAddr, testAdminThirdAddr}
	suite.kvsKeeper.SetAcl(suite.goCtx, types.Acl{
		Adresses: testAddresses,
	})

	msg := &types.MsgDataConfirmation{
		Creator: testAdminFirstAddr,
		Key:     testKey,
	}
	_, err := suite.msgServer.DataConfirmation(ctx, msg)
	if err == nil {
		suite.T().Fail()
	}
}

func (suite *KvsTestSuite) TestKvs_DataConfirmation_AddressInvalid_Fail() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testKey := fmt.Sprint("testKey%i", rand.Int())
	testValue := fmt.Sprint("testKey%i", rand.Int())

	testAddresses := []string{testAdminFirstAddr, testAdminSecondAddr, testAdminThirdAddr}
	suite.kvsKeeper.SetAcl(suite.goCtx, types.Acl{
		Adresses: testAddresses,
	})

	suite.kvsKeeper.SetProposal(suite.goCtx, types.Proposal{
		Index:           testKey,
		Value:           testValue,
		Acknowledgments: []string{},
	})

	msg := &types.MsgDataConfirmation{
		Creator: testCreatorAddr,
		Key:     testKey,
	}
	_, err := suite.msgServer.DataConfirmation(ctx, msg)
	if err == nil {
		suite.T().Fail()
	}
}

func (suite *KvsTestSuite) TestKvs_DataConfirmation_AddressNotRegistred_Fail() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testKey := fmt.Sprint("testKey%i", rand.Int())
	testValue := fmt.Sprint("testKey%i", rand.Int())

	suite.kvsKeeper.SetProposal(suite.goCtx, types.Proposal{
		Index:           testKey,
		Value:           testValue,
		Acknowledgments: []string{},
	})

	msg := &types.MsgDataConfirmation{
		Creator: testAdminFirstAddr,
		Key:     testKey,
	}
	_, err := suite.msgServer.DataConfirmation(ctx, msg)
	if err == nil {
		suite.T().Fail()
	}
}

func (suite *KvsTestSuite) TestKvs_DataConfirmation_AllConfirmation_Success() {
	ctx := sdk.WrapSDKContext(suite.goCtx)
	testKey := fmt.Sprint("testKey%i", rand.Int())
	testValue := fmt.Sprint("testKey%i", rand.Int())

	testAddresses := []string{testAdminFirstAddr, testAdminSecondAddr, testAdminThirdAddr}
	suite.kvsKeeper.SetAcl(suite.goCtx, types.Acl{
		Adresses: testAddresses,
	})

	suite.kvsKeeper.SetProposal(suite.goCtx, types.Proposal{
		Index:           testKey,
		Value:           testValue,
		Acknowledgments: []string{},
	})

	msg := &types.MsgDataConfirmation{
		Creator: testAdminFirstAddr,
		Key:     testKey,
	}
	_, err := suite.msgServer.DataConfirmation(ctx, msg)
	if err != nil {
		suite.T().Fail()
	}
	msg = &types.MsgDataConfirmation{
		Creator: testAdminSecondAddr,
		Key:     testKey,
	}
	_, err = suite.msgServer.DataConfirmation(ctx, msg)
	if err != nil {
		suite.T().Fail()
	}
	msg = &types.MsgDataConfirmation{
		Creator: testAdminThirdAddr,
		Key:     testKey,
	}
	_, err = suite.msgServer.DataConfirmation(ctx, msg)
	if err != nil {
		suite.T().Fail()
	}
	testAck := []string{testAdminFirstAddr, testAdminSecondAddr, testAdminThirdAddr}
	prop, exuist := suite.kvsKeeper.GetProposal(suite.goCtx, testKey)
	assert.Equal(suite.T(), true, exuist)
	assert.Equal(suite.T(), testKey, prop.Index)
	assert.Equal(suite.T(), testValue, prop.Value)
	assert.Equal(suite.T(), testAck, prop.Acknowledgments)
	data, exuist := suite.kvsKeeper.GetData(suite.goCtx, testKey)
	assert.Equal(suite.T(), true, exuist)
	assert.Equal(suite.T(), testKey, data.Index)
	assert.Equal(suite.T(), testValue, data.Value)
}
