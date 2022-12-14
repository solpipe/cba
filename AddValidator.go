// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package solmate_cba

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/SolmateDev/solana-go"
	ag_format "github.com/SolmateDev/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Create a ValidatorMember account.
//
// # Panics
//
// Panics if .
//
// # Errors
//
// This function will return an error if the bump is not correct.
type AddValidator struct {

	// [0] = [] controller
	//
	// [1] = [WRITE] validatorManager
	//
	// [2] = [] vote
	//
	// [3] = [] stake
	//
	// [4] = [SIGNER] voteAdmin
	//
	// [5] = [WRITE, SIGNER] validatorAdmin
	//
	// [6] = [] systemProgram
	//
	// [7] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewAddValidatorInstructionBuilder creates a new `AddValidator` instruction builder.
func NewAddValidatorInstructionBuilder() *AddValidator {
	nd := &AddValidator{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetControllerAccount sets the "controller" account.
func (inst *AddValidator) SetControllerAccount(controller ag_solanago.PublicKey) *AddValidator {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(controller)
	return inst
}

// GetControllerAccount gets the "controller" account.
func (inst *AddValidator) GetControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetValidatorManagerAccount sets the "validatorManager" account.
func (inst *AddValidator) SetValidatorManagerAccount(validatorManager ag_solanago.PublicKey) *AddValidator {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(validatorManager).WRITE()
	return inst
}

// GetValidatorManagerAccount gets the "validatorManager" account.
func (inst *AddValidator) GetValidatorManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetVoteAccount sets the "vote" account.
func (inst *AddValidator) SetVoteAccount(vote ag_solanago.PublicKey) *AddValidator {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(vote)
	return inst
}

// GetVoteAccount gets the "vote" account.
func (inst *AddValidator) GetVoteAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetStakeAccount sets the "stake" account.
func (inst *AddValidator) SetStakeAccount(stake ag_solanago.PublicKey) *AddValidator {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(stake)
	return inst
}

// GetStakeAccount gets the "stake" account.
func (inst *AddValidator) GetStakeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetVoteAdminAccount sets the "voteAdmin" account.
func (inst *AddValidator) SetVoteAdminAccount(voteAdmin ag_solanago.PublicKey) *AddValidator {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(voteAdmin).SIGNER()
	return inst
}

// GetVoteAdminAccount gets the "voteAdmin" account.
func (inst *AddValidator) GetVoteAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetValidatorAdminAccount sets the "validatorAdmin" account.
func (inst *AddValidator) SetValidatorAdminAccount(validatorAdmin ag_solanago.PublicKey) *AddValidator {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(validatorAdmin).WRITE().SIGNER()
	return inst
}

// GetValidatorAdminAccount gets the "validatorAdmin" account.
func (inst *AddValidator) GetValidatorAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *AddValidator) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *AddValidator {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *AddValidator) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetRentAccount sets the "rent" account.
func (inst *AddValidator) SetRentAccount(rent ag_solanago.PublicKey) *AddValidator {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *AddValidator) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

func (inst AddValidator) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_AddValidator,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst AddValidator) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *AddValidator) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Controller is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ValidatorManager is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Vote is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Stake is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.VoteAdmin is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ValidatorAdmin is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *AddValidator) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("AddValidator")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("      controller", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("validatorManager", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("            vote", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           stake", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       voteAdmin", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("  validatorAdmin", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("   systemProgram", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("            rent", inst.AccountMetaSlice.Get(7)))
					})
				})
		})
}

func (obj AddValidator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *AddValidator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewAddValidatorInstruction declares a new AddValidator instruction with the provided parameters and accounts.
func NewAddValidatorInstruction(
	// Accounts:
	controller ag_solanago.PublicKey,
	validatorManager ag_solanago.PublicKey,
	vote ag_solanago.PublicKey,
	stake ag_solanago.PublicKey,
	voteAdmin ag_solanago.PublicKey,
	validatorAdmin ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *AddValidator {
	return NewAddValidatorInstructionBuilder().
		SetControllerAccount(controller).
		SetValidatorManagerAccount(validatorManager).
		SetVoteAccount(vote).
		SetStakeAccount(stake).
		SetVoteAdminAccount(voteAdmin).
		SetValidatorAdminAccount(validatorAdmin).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent)
}
