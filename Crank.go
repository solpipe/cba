// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package solmate_cba

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/SolmateDev/solana-go"
	ag_format "github.com/SolmateDev/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Propogate periods forward in time and distribute revenue to stakeholders.
// Make sure to pick the right payout that will match the period that will be cranked (new period has started, but has not been cranked yet).
// The crank instruction is the only way to change the bandwidth allocations!
// # Errors
//
// This function will return an error if .
type Crank struct {

	// [0] = [] controller
	//
	// [1] = [] pipeline
	//
	// [2] = [WRITE] pcVault
	//
	// [3] = [WRITE] bids
	//
	// [4] = [WRITE] periods
	//
	// [5] = [] payout
	//
	// [6] = [WRITE] crankerFund
	//
	// [7] = [WRITE, SIGNER] cranker
	//
	// [8] = [] tokenProgram
	//
	// [9] = [] clock
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewCrankInstructionBuilder creates a new `Crank` instruction builder.
func NewCrankInstructionBuilder() *Crank {
	nd := &Crank{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetControllerAccount sets the "controller" account.
func (inst *Crank) SetControllerAccount(controller ag_solanago.PublicKey) *Crank {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(controller)
	return inst
}

// GetControllerAccount gets the "controller" account.
func (inst *Crank) GetControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPipelineAccount sets the "pipeline" account.
func (inst *Crank) SetPipelineAccount(pipeline ag_solanago.PublicKey) *Crank {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(pipeline)
	return inst
}

// GetPipelineAccount gets the "pipeline" account.
func (inst *Crank) GetPipelineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPcVaultAccount sets the "pcVault" account.
func (inst *Crank) SetPcVaultAccount(pcVault ag_solanago.PublicKey) *Crank {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(pcVault).WRITE()
	return inst
}

// GetPcVaultAccount gets the "pcVault" account.
func (inst *Crank) GetPcVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetBidsAccount sets the "bids" account.
func (inst *Crank) SetBidsAccount(bids ag_solanago.PublicKey) *Crank {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(bids).WRITE()
	return inst
}

// GetBidsAccount gets the "bids" account.
func (inst *Crank) GetBidsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPeriodsAccount sets the "periods" account.
func (inst *Crank) SetPeriodsAccount(periods ag_solanago.PublicKey) *Crank {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(periods).WRITE()
	return inst
}

// GetPeriodsAccount gets the "periods" account.
func (inst *Crank) GetPeriodsAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPayoutAccount sets the "payout" account.
func (inst *Crank) SetPayoutAccount(payout ag_solanago.PublicKey) *Crank {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(payout)
	return inst
}

// GetPayoutAccount gets the "payout" account.
func (inst *Crank) GetPayoutAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetCrankerFundAccount sets the "crankerFund" account.
func (inst *Crank) SetCrankerFundAccount(crankerFund ag_solanago.PublicKey) *Crank {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(crankerFund).WRITE()
	return inst
}

// GetCrankerFundAccount gets the "crankerFund" account.
func (inst *Crank) GetCrankerFundAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetCrankerAccount sets the "cranker" account.
func (inst *Crank) SetCrankerAccount(cranker ag_solanago.PublicKey) *Crank {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(cranker).WRITE().SIGNER()
	return inst
}

// GetCrankerAccount gets the "cranker" account.
func (inst *Crank) GetCrankerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *Crank) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Crank {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *Crank) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetClockAccount sets the "clock" account.
func (inst *Crank) SetClockAccount(clock ag_solanago.PublicKey) *Crank {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *Crank) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst Crank) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Crank,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Crank) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Crank) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Controller is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Pipeline is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PcVault is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Bids is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Periods is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Payout is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.CrankerFund is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Cranker is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Clock is not set")
		}
	}
	return nil
}

func (inst *Crank) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Crank")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("  controller", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("    pipeline", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     pcVault", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("        bids", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     periods", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("      payout", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta(" crankerFund", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("     cranker", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("tokenProgram", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("       clock", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj Crank) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *Crank) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewCrankInstruction declares a new Crank instruction with the provided parameters and accounts.
func NewCrankInstruction(
	// Accounts:
	controller ag_solanago.PublicKey,
	pipeline ag_solanago.PublicKey,
	pcVault ag_solanago.PublicKey,
	bids ag_solanago.PublicKey,
	periods ag_solanago.PublicKey,
	payout ag_solanago.PublicKey,
	crankerFund ag_solanago.PublicKey,
	cranker ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	clock ag_solanago.PublicKey) *Crank {
	return NewCrankInstructionBuilder().
		SetControllerAccount(controller).
		SetPipelineAccount(pipeline).
		SetPcVaultAccount(pcVault).
		SetBidsAccount(bids).
		SetPeriodsAccount(periods).
		SetPayoutAccount(payout).
		SetCrankerFundAccount(crankerFund).
		SetCrankerAccount(cranker).
		SetTokenProgramAccount(tokenProgram).
		SetClockAccount(clock)
}
