// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package solmate_cba

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/SolmateDev/solana-go"
)

type BidReceipt struct {
	User             ag_solanago.PublicKey
	Payout           ag_solanago.PublicKey
	Bump             uint8
	LastTx           ag_solanago.Hash
	TxSentMerkleroot ag_solanago.Hash
	TxSent           uint32
}

var BidReceiptDiscriminator = [8]byte{186, 150, 141, 135, 59, 122, 39, 99}

func (obj BidReceipt) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(BidReceiptDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `User` param:
	err = encoder.Encode(obj.User)
	if err != nil {
		return err
	}
	// Serialize `Payout` param:
	err = encoder.Encode(obj.Payout)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `LastTx` param:
	err = encoder.Encode(obj.LastTx)
	if err != nil {
		return err
	}
	// Serialize `TxSentMerkleroot` param:
	err = encoder.Encode(obj.TxSentMerkleroot)
	if err != nil {
		return err
	}
	// Serialize `TxSent` param:
	err = encoder.Encode(obj.TxSent)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BidReceipt) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(BidReceiptDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[186 150 141 135 59 122 39 99]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `User`:
	err = decoder.Decode(&obj.User)
	if err != nil {
		return err
	}
	// Deserialize `Payout`:
	err = decoder.Decode(&obj.Payout)
	if err != nil {
		return err
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `LastTx`:
	err = decoder.Decode(&obj.LastTx)
	if err != nil {
		return err
	}
	// Deserialize `TxSentMerkleroot`:
	err = decoder.Decode(&obj.TxSentMerkleroot)
	if err != nil {
		return err
	}
	// Deserialize `TxSent`:
	err = decoder.Decode(&obj.TxSent)
	if err != nil {
		return err
	}
	return nil
}

type BidList struct {
	Pipeline             ag_solanago.PublicKey
	Book                 []Bid
	LastPeriodStart      uint64
	BandwidthDenominator uint64
	TotalDeposits        uint64
}

var BidListDiscriminator = [8]byte{233, 127, 13, 29, 123, 209, 192, 79}

func (obj BidList) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(BidListDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pipeline` param:
	err = encoder.Encode(obj.Pipeline)
	if err != nil {
		return err
	}
	// Serialize `Book` param:
	err = encoder.Encode(obj.Book)
	if err != nil {
		return err
	}
	// Serialize `LastPeriodStart` param:
	err = encoder.Encode(obj.LastPeriodStart)
	if err != nil {
		return err
	}
	// Serialize `BandwidthDenominator` param:
	err = encoder.Encode(obj.BandwidthDenominator)
	if err != nil {
		return err
	}
	// Serialize `TotalDeposits` param:
	err = encoder.Encode(obj.TotalDeposits)
	if err != nil {
		return err
	}
	return nil
}

func (obj *BidList) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(BidListDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[233 127 13 29 123 209 192 79]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pipeline`:
	err = decoder.Decode(&obj.Pipeline)
	if err != nil {
		return err
	}
	// Deserialize `Book`:
	err = decoder.Decode(&obj.Book)
	if err != nil {
		return err
	}
	// Deserialize `LastPeriodStart`:
	err = decoder.Decode(&obj.LastPeriodStart)
	if err != nil {
		return err
	}
	// Deserialize `BandwidthDenominator`:
	err = decoder.Decode(&obj.BandwidthDenominator)
	if err != nil {
		return err
	}
	// Deserialize `TotalDeposits`:
	err = decoder.Decode(&obj.TotalDeposits)
	if err != nil {
		return err
	}
	return nil
}

type Controller struct {
	ControllerBump         uint8
	PcVaultBump            uint8
	Admin                  ag_solanago.PublicKey
	CrankAuthority         ag_solanago.PublicKey
	ControllerFee          [2]uint64
	PcVault                ag_solanago.PublicKey
	PcMint                 ag_solanago.PublicKey
	PcVaultAllocatedAmount uint64
}

var ControllerDiscriminator = [8]byte{184, 79, 171, 0, 183, 43, 113, 110}

func (obj Controller) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ControllerDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `ControllerBump` param:
	err = encoder.Encode(obj.ControllerBump)
	if err != nil {
		return err
	}
	// Serialize `PcVaultBump` param:
	err = encoder.Encode(obj.PcVaultBump)
	if err != nil {
		return err
	}
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	// Serialize `CrankAuthority` param:
	err = encoder.Encode(obj.CrankAuthority)
	if err != nil {
		return err
	}
	// Serialize `ControllerFee` param:
	err = encoder.Encode(obj.ControllerFee)
	if err != nil {
		return err
	}
	// Serialize `PcVault` param:
	err = encoder.Encode(obj.PcVault)
	if err != nil {
		return err
	}
	// Serialize `PcMint` param:
	err = encoder.Encode(obj.PcMint)
	if err != nil {
		return err
	}
	// Serialize `PcVaultAllocatedAmount` param:
	err = encoder.Encode(obj.PcVaultAllocatedAmount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Controller) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ControllerDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[184 79 171 0 183 43 113 110]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `ControllerBump`:
	err = decoder.Decode(&obj.ControllerBump)
	if err != nil {
		return err
	}
	// Deserialize `PcVaultBump`:
	err = decoder.Decode(&obj.PcVaultBump)
	if err != nil {
		return err
	}
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	// Deserialize `CrankAuthority`:
	err = decoder.Decode(&obj.CrankAuthority)
	if err != nil {
		return err
	}
	// Deserialize `ControllerFee`:
	err = decoder.Decode(&obj.ControllerFee)
	if err != nil {
		return err
	}
	// Deserialize `PcVault`:
	err = decoder.Decode(&obj.PcVault)
	if err != nil {
		return err
	}
	// Deserialize `PcMint`:
	err = decoder.Decode(&obj.PcMint)
	if err != nil {
		return err
	}
	// Deserialize `PcVaultAllocatedAmount`:
	err = decoder.Decode(&obj.PcVaultAllocatedAmount)
	if err != nil {
		return err
	}
	return nil
}

type Payout struct {
	Pipeline             ag_solanago.PublicKey
	Period               Period
	ControllerFee        [2]uint64
	DecayRate            [2]uint64
	CrankFee             [2]uint64
	ValidatorPayoutShare [2]uint64
	Revenue              uint64
	ControllerRevenue    uint64
	ValidatorRevenue     uint64
	ValidatorWithdrawn   uint64
	StakerRevenue        uint64
	StakerWithdrawn      uint64
	StakerCount          uint32
	TotalTxCount         uint32
	TotalStake           uint64
	ValidatorCount       uint32
}

var PayoutDiscriminator = [8]byte{69, 45, 245, 131, 218, 101, 158, 228}

func (obj Payout) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PayoutDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pipeline` param:
	err = encoder.Encode(obj.Pipeline)
	if err != nil {
		return err
	}
	// Serialize `Period` param:
	err = encoder.Encode(obj.Period)
	if err != nil {
		return err
	}
	// Serialize `ControllerFee` param:
	err = encoder.Encode(obj.ControllerFee)
	if err != nil {
		return err
	}
	// Serialize `DecayRate` param:
	err = encoder.Encode(obj.DecayRate)
	if err != nil {
		return err
	}
	// Serialize `CrankFee` param:
	err = encoder.Encode(obj.CrankFee)
	if err != nil {
		return err
	}
	// Serialize `ValidatorPayoutShare` param:
	err = encoder.Encode(obj.ValidatorPayoutShare)
	if err != nil {
		return err
	}
	// Serialize `Revenue` param:
	err = encoder.Encode(obj.Revenue)
	if err != nil {
		return err
	}
	// Serialize `ControllerRevenue` param:
	err = encoder.Encode(obj.ControllerRevenue)
	if err != nil {
		return err
	}
	// Serialize `ValidatorRevenue` param:
	err = encoder.Encode(obj.ValidatorRevenue)
	if err != nil {
		return err
	}
	// Serialize `ValidatorWithdrawn` param:
	err = encoder.Encode(obj.ValidatorWithdrawn)
	if err != nil {
		return err
	}
	// Serialize `StakerRevenue` param:
	err = encoder.Encode(obj.StakerRevenue)
	if err != nil {
		return err
	}
	// Serialize `StakerWithdrawn` param:
	err = encoder.Encode(obj.StakerWithdrawn)
	if err != nil {
		return err
	}
	// Serialize `StakerCount` param:
	err = encoder.Encode(obj.StakerCount)
	if err != nil {
		return err
	}
	// Serialize `TotalTxCount` param:
	err = encoder.Encode(obj.TotalTxCount)
	if err != nil {
		return err
	}
	// Serialize `TotalStake` param:
	err = encoder.Encode(obj.TotalStake)
	if err != nil {
		return err
	}
	// Serialize `ValidatorCount` param:
	err = encoder.Encode(obj.ValidatorCount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Payout) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PayoutDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[69 45 245 131 218 101 158 228]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pipeline`:
	err = decoder.Decode(&obj.Pipeline)
	if err != nil {
		return err
	}
	// Deserialize `Period`:
	err = decoder.Decode(&obj.Period)
	if err != nil {
		return err
	}
	// Deserialize `ControllerFee`:
	err = decoder.Decode(&obj.ControllerFee)
	if err != nil {
		return err
	}
	// Deserialize `DecayRate`:
	err = decoder.Decode(&obj.DecayRate)
	if err != nil {
		return err
	}
	// Deserialize `CrankFee`:
	err = decoder.Decode(&obj.CrankFee)
	if err != nil {
		return err
	}
	// Deserialize `ValidatorPayoutShare`:
	err = decoder.Decode(&obj.ValidatorPayoutShare)
	if err != nil {
		return err
	}
	// Deserialize `Revenue`:
	err = decoder.Decode(&obj.Revenue)
	if err != nil {
		return err
	}
	// Deserialize `ControllerRevenue`:
	err = decoder.Decode(&obj.ControllerRevenue)
	if err != nil {
		return err
	}
	// Deserialize `ValidatorRevenue`:
	err = decoder.Decode(&obj.ValidatorRevenue)
	if err != nil {
		return err
	}
	// Deserialize `ValidatorWithdrawn`:
	err = decoder.Decode(&obj.ValidatorWithdrawn)
	if err != nil {
		return err
	}
	// Deserialize `StakerRevenue`:
	err = decoder.Decode(&obj.StakerRevenue)
	if err != nil {
		return err
	}
	// Deserialize `StakerWithdrawn`:
	err = decoder.Decode(&obj.StakerWithdrawn)
	if err != nil {
		return err
	}
	// Deserialize `StakerCount`:
	err = decoder.Decode(&obj.StakerCount)
	if err != nil {
		return err
	}
	// Deserialize `TotalTxCount`:
	err = decoder.Decode(&obj.TotalTxCount)
	if err != nil {
		return err
	}
	// Deserialize `TotalStake`:
	err = decoder.Decode(&obj.TotalStake)
	if err != nil {
		return err
	}
	// Deserialize `ValidatorCount`:
	err = decoder.Decode(&obj.ValidatorCount)
	if err != nil {
		return err
	}
	return nil
}

type PeriodRing struct {
	Pipeline ag_solanago.PublicKey
	Ring     []PeriodWithPayout
	Start    uint16
	Length   uint16
}

var PeriodRingDiscriminator = [8]byte{61, 191, 59, 143, 226, 235, 104, 26}

func (obj PeriodRing) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PeriodRingDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Pipeline` param:
	err = encoder.Encode(obj.Pipeline)
	if err != nil {
		return err
	}
	// Serialize `Ring` param:
	err = encoder.Encode(obj.Ring)
	if err != nil {
		return err
	}
	// Serialize `Start` param:
	err = encoder.Encode(obj.Start)
	if err != nil {
		return err
	}
	// Serialize `Length` param:
	err = encoder.Encode(obj.Length)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PeriodRing) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PeriodRingDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[61 191 59 143 226 235 104 26]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Pipeline`:
	err = decoder.Decode(&obj.Pipeline)
	if err != nil {
		return err
	}
	// Deserialize `Ring`:
	err = decoder.Decode(&obj.Ring)
	if err != nil {
		return err
	}
	// Deserialize `Start`:
	err = decoder.Decode(&obj.Start)
	if err != nil {
		return err
	}
	// Deserialize `Length`:
	err = decoder.Decode(&obj.Length)
	if err != nil {
		return err
	}
	return nil
}

type Pipeline struct {
	Controller           ag_solanago.PublicKey
	PcVault              ag_solanago.PublicKey
	PcVaultBump          uint8
	Admin                ag_solanago.PublicKey
	CrankFeeRate         [2]uint64
	Periods              ag_solanago.PublicKey
	Bids                 ag_solanago.PublicKey
	DecayRate            [2]uint64
	ValidatorPayoutShare [2]uint64
	BandwidthAllotment   uint16
	WithdrawBalance      uint64
}

var PipelineDiscriminator = [8]byte{30, 82, 16, 218, 196, 77, 115, 224}

func (obj Pipeline) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(PipelineDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Controller` param:
	err = encoder.Encode(obj.Controller)
	if err != nil {
		return err
	}
	// Serialize `PcVault` param:
	err = encoder.Encode(obj.PcVault)
	if err != nil {
		return err
	}
	// Serialize `PcVaultBump` param:
	err = encoder.Encode(obj.PcVaultBump)
	if err != nil {
		return err
	}
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	// Serialize `CrankFeeRate` param:
	err = encoder.Encode(obj.CrankFeeRate)
	if err != nil {
		return err
	}
	// Serialize `Periods` param:
	err = encoder.Encode(obj.Periods)
	if err != nil {
		return err
	}
	// Serialize `Bids` param:
	err = encoder.Encode(obj.Bids)
	if err != nil {
		return err
	}
	// Serialize `DecayRate` param:
	err = encoder.Encode(obj.DecayRate)
	if err != nil {
		return err
	}
	// Serialize `ValidatorPayoutShare` param:
	err = encoder.Encode(obj.ValidatorPayoutShare)
	if err != nil {
		return err
	}
	// Serialize `BandwidthAllotment` param:
	err = encoder.Encode(obj.BandwidthAllotment)
	if err != nil {
		return err
	}
	// Serialize `WithdrawBalance` param:
	err = encoder.Encode(obj.WithdrawBalance)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Pipeline) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(PipelineDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[30 82 16 218 196 77 115 224]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Controller`:
	err = decoder.Decode(&obj.Controller)
	if err != nil {
		return err
	}
	// Deserialize `PcVault`:
	err = decoder.Decode(&obj.PcVault)
	if err != nil {
		return err
	}
	// Deserialize `PcVaultBump`:
	err = decoder.Decode(&obj.PcVaultBump)
	if err != nil {
		return err
	}
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	// Deserialize `CrankFeeRate`:
	err = decoder.Decode(&obj.CrankFeeRate)
	if err != nil {
		return err
	}
	// Deserialize `Periods`:
	err = decoder.Decode(&obj.Periods)
	if err != nil {
		return err
	}
	// Deserialize `Bids`:
	err = decoder.Decode(&obj.Bids)
	if err != nil {
		return err
	}
	// Deserialize `DecayRate`:
	err = decoder.Decode(&obj.DecayRate)
	if err != nil {
		return err
	}
	// Deserialize `ValidatorPayoutShare`:
	err = decoder.Decode(&obj.ValidatorPayoutShare)
	if err != nil {
		return err
	}
	// Deserialize `BandwidthAllotment`:
	err = decoder.Decode(&obj.BandwidthAllotment)
	if err != nil {
		return err
	}
	// Deserialize `WithdrawBalance`:
	err = decoder.Decode(&obj.WithdrawBalance)
	if err != nil {
		return err
	}
	return nil
}

type Receipt struct {
	Payout           ag_solanago.PublicKey
	Validator        ag_solanago.PublicKey
	MsgCounter       uint32
	Stake            uint64
	StakerCounter    uint32
	TxSentMerkleroot ag_solanago.Hash
	TxSent           uint32
	LastTx           ag_solanago.Hash
}

var ReceiptDiscriminator = [8]byte{39, 154, 73, 106, 80, 102, 145, 153}

func (obj Receipt) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ReceiptDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Payout` param:
	err = encoder.Encode(obj.Payout)
	if err != nil {
		return err
	}
	// Serialize `Validator` param:
	err = encoder.Encode(obj.Validator)
	if err != nil {
		return err
	}
	// Serialize `MsgCounter` param:
	err = encoder.Encode(obj.MsgCounter)
	if err != nil {
		return err
	}
	// Serialize `Stake` param:
	err = encoder.Encode(obj.Stake)
	if err != nil {
		return err
	}
	// Serialize `StakerCounter` param:
	err = encoder.Encode(obj.StakerCounter)
	if err != nil {
		return err
	}
	// Serialize `TxSentMerkleroot` param:
	err = encoder.Encode(obj.TxSentMerkleroot)
	if err != nil {
		return err
	}
	// Serialize `TxSent` param:
	err = encoder.Encode(obj.TxSent)
	if err != nil {
		return err
	}
	// Serialize `LastTx` param:
	err = encoder.Encode(obj.LastTx)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Receipt) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ReceiptDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[39 154 73 106 80 102 145 153]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Payout`:
	err = decoder.Decode(&obj.Payout)
	if err != nil {
		return err
	}
	// Deserialize `Validator`:
	err = decoder.Decode(&obj.Validator)
	if err != nil {
		return err
	}
	// Deserialize `MsgCounter`:
	err = decoder.Decode(&obj.MsgCounter)
	if err != nil {
		return err
	}
	// Deserialize `Stake`:
	err = decoder.Decode(&obj.Stake)
	if err != nil {
		return err
	}
	// Deserialize `StakerCounter`:
	err = decoder.Decode(&obj.StakerCounter)
	if err != nil {
		return err
	}
	// Deserialize `TxSentMerkleroot`:
	err = decoder.Decode(&obj.TxSentMerkleroot)
	if err != nil {
		return err
	}
	// Deserialize `TxSent`:
	err = decoder.Decode(&obj.TxSent)
	if err != nil {
		return err
	}
	// Deserialize `LastTx`:
	err = decoder.Decode(&obj.LastTx)
	if err != nil {
		return err
	}
	return nil
}

type StakerMember struct {
	Bump           uint8
	Receipt        ag_solanago.PublicKey
	Stake          ag_solanago.PublicKey
	PcVault        ag_solanago.PublicKey
	DelegatedStake uint64
}

var StakerMemberDiscriminator = [8]byte{21, 191, 126, 39, 175, 235, 92, 31}

func (obj StakerMember) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(StakerMemberDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `Receipt` param:
	err = encoder.Encode(obj.Receipt)
	if err != nil {
		return err
	}
	// Serialize `Stake` param:
	err = encoder.Encode(obj.Stake)
	if err != nil {
		return err
	}
	// Serialize `PcVault` param:
	err = encoder.Encode(obj.PcVault)
	if err != nil {
		return err
	}
	// Serialize `DelegatedStake` param:
	err = encoder.Encode(obj.DelegatedStake)
	if err != nil {
		return err
	}
	return nil
}

func (obj *StakerMember) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(StakerMemberDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[21 191 126 39 175 235 92 31]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `Receipt`:
	err = decoder.Decode(&obj.Receipt)
	if err != nil {
		return err
	}
	// Deserialize `Stake`:
	err = decoder.Decode(&obj.Stake)
	if err != nil {
		return err
	}
	// Deserialize `PcVault`:
	err = decoder.Decode(&obj.PcVault)
	if err != nil {
		return err
	}
	// Deserialize `DelegatedStake`:
	err = decoder.Decode(&obj.DelegatedStake)
	if err != nil {
		return err
	}
	return nil
}

type ValidatorMember struct {
	Bump                  uint8
	Receipt               ag_solanago.PublicKey
	Vote                  ag_solanago.PublicKey
	Admin                 ag_solanago.PublicKey
	Controller            ag_solanago.PublicKey
	IsSet                 bool
	HasValidatorWithdrawn bool
}

var ValidatorMemberDiscriminator = [8]byte{176, 174, 147, 68, 122, 100, 104, 120}

func (obj ValidatorMember) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ValidatorMemberDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Bump` param:
	err = encoder.Encode(obj.Bump)
	if err != nil {
		return err
	}
	// Serialize `Receipt` param:
	err = encoder.Encode(obj.Receipt)
	if err != nil {
		return err
	}
	// Serialize `Vote` param:
	err = encoder.Encode(obj.Vote)
	if err != nil {
		return err
	}
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	// Serialize `Controller` param:
	err = encoder.Encode(obj.Controller)
	if err != nil {
		return err
	}
	// Serialize `IsSet` param:
	err = encoder.Encode(obj.IsSet)
	if err != nil {
		return err
	}
	// Serialize `HasValidatorWithdrawn` param:
	err = encoder.Encode(obj.HasValidatorWithdrawn)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ValidatorMember) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ValidatorMemberDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[176 174 147 68 122 100 104 120]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Bump`:
	err = decoder.Decode(&obj.Bump)
	if err != nil {
		return err
	}
	// Deserialize `Receipt`:
	err = decoder.Decode(&obj.Receipt)
	if err != nil {
		return err
	}
	// Deserialize `Vote`:
	err = decoder.Decode(&obj.Vote)
	if err != nil {
		return err
	}
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	// Deserialize `Controller`:
	err = decoder.Decode(&obj.Controller)
	if err != nil {
		return err
	}
	// Deserialize `IsSet`:
	err = decoder.Decode(&obj.IsSet)
	if err != nil {
		return err
	}
	// Deserialize `HasValidatorWithdrawn`:
	err = decoder.Decode(&obj.HasValidatorWithdrawn)
	if err != nil {
		return err
	}
	return nil
}