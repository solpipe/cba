// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package solmate_cba

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/SolmateDev/solana-go"
	ag_format "github.com/SolmateDev/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// SetValidator is the `setValidator` instruction.
type SetValidator struct {

	// [0] = [] controller
	//
	// [1] = [WRITE] pipeline
	//
	// [2] = [] payout
	//
	// [3] = [] validatorMember
	//
	// [4] = [WRITE, SIGNER] receipt
	//
	// [5] = [WRITE, SIGNER] validatorAdmin
	//
	// [6] = [] clock
	//
	// [7] = [] systemProgram
	//
	// [8] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewSetValidatorInstructionBuilder creates a new `SetValidator` instruction builder.
func NewSetValidatorInstructionBuilder() *SetValidator {
	nd := &SetValidator{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetControllerAccount sets the "controller" account.
func (inst *SetValidator) SetControllerAccount(controller ag_solanago.PublicKey) *SetValidator {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(controller)
	return inst
}

// GetControllerAccount gets the "controller" account.
func (inst *SetValidator) GetControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetPipelineAccount sets the "pipeline" account.
func (inst *SetValidator) SetPipelineAccount(pipeline ag_solanago.PublicKey) *SetValidator {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(pipeline).WRITE()
	return inst
}

// GetPipelineAccount gets the "pipeline" account.
func (inst *SetValidator) GetPipelineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayoutAccount sets the "payout" account.
func (inst *SetValidator) SetPayoutAccount(payout ag_solanago.PublicKey) *SetValidator {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payout)
	return inst
}

// GetPayoutAccount gets the "payout" account.
func (inst *SetValidator) GetPayoutAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetValidatorMemberAccount sets the "validatorMember" account.
func (inst *SetValidator) SetValidatorMemberAccount(validatorMember ag_solanago.PublicKey) *SetValidator {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(validatorMember)
	return inst
}

// GetValidatorMemberAccount gets the "validatorMember" account.
func (inst *SetValidator) GetValidatorMemberAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetReceiptAccount sets the "receipt" account.
func (inst *SetValidator) SetReceiptAccount(receipt ag_solanago.PublicKey) *SetValidator {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(receipt).WRITE().SIGNER()
	return inst
}

// GetReceiptAccount gets the "receipt" account.
func (inst *SetValidator) GetReceiptAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetValidatorAdminAccount sets the "validatorAdmin" account.
func (inst *SetValidator) SetValidatorAdminAccount(validatorAdmin ag_solanago.PublicKey) *SetValidator {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(validatorAdmin).WRITE().SIGNER()
	return inst
}

// GetValidatorAdminAccount gets the "validatorAdmin" account.
func (inst *SetValidator) GetValidatorAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetClockAccount sets the "clock" account.
func (inst *SetValidator) SetClockAccount(clock ag_solanago.PublicKey) *SetValidator {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *SetValidator) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *SetValidator) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *SetValidator {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *SetValidator) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetRentAccount sets the "rent" account.
func (inst *SetValidator) SetRentAccount(rent ag_solanago.PublicKey) *SetValidator {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *SetValidator) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst SetValidator) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_SetValidator,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst SetValidator) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *SetValidator) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Controller is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Pipeline is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payout is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ValidatorMember is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Receipt is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ValidatorAdmin is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Clock is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *SetValidator) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("SetValidator")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("     controller", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       pipeline", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("         payout", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("validatorMember", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("        receipt", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta(" validatorAdmin", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("          clock", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("  systemProgram", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("           rent", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj SetValidator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *SetValidator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewSetValidatorInstruction declares a new SetValidator instruction with the provided parameters and accounts.
func NewSetValidatorInstruction(
	// Accounts:
	controller ag_solanago.PublicKey,
	pipeline ag_solanago.PublicKey,
	payout ag_solanago.PublicKey,
	validatorMember ag_solanago.PublicKey,
	receipt ag_solanago.PublicKey,
	validatorAdmin ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *SetValidator {
	return NewSetValidatorInstructionBuilder().
		SetControllerAccount(controller).
		SetPipelineAccount(pipeline).
		SetPayoutAccount(payout).
		SetValidatorMemberAccount(validatorMember).
		SetReceiptAccount(receipt).
		SetValidatorAdminAccount(validatorAdmin).
		SetClockAccount(clock).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
