// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package EME

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// Encrypted Message Envelope
type EME struct {
	_tab flatbuffers.Table
}

func GetRootAsEME(buf []byte, offset flatbuffers.UOffsetT) *EME {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &EME{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsEME(buf []byte, offset flatbuffers.UOffsetT) *EME {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &EME{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *EME) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *EME) Table() flatbuffers.Table {
	return rcv._tab
}

/// Encrypted data blob, containing the ciphertext of the original plaintext message.
func (rcv *EME) ENCRYPTED_BLOB(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *EME) ENCRYPTED_BLOBLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *EME) ENCRYPTED_BLOBBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Encrypted data blob, containing the ciphertext of the original plaintext message.
func (rcv *EME) MutateENCRYPTED_BLOB(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

/// Temporary public key used for the encryption session, contributing to the derivation of the shared secret.
func (rcv *EME) EPHEMERAL_PUBLIC_KEY() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Temporary public key used for the encryption session, contributing to the derivation of the shared secret.
/// Message Authentication Code to verify the integrity and authenticity of the encrypted message.
func (rcv *EME) MAC() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Message Authentication Code to verify the integrity and authenticity of the encrypted message.
/// Unique value used to ensure that the same plaintext produces a different ciphertext for each encryption.
func (rcv *EME) NONCE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Unique value used to ensure that the same plaintext produces a different ciphertext for each encryption.
/// Additional authentication tag used in some encryption schemes for integrity and authenticity verification.
func (rcv *EME) TAG() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Additional authentication tag used in some encryption schemes for integrity and authenticity verification.
/// Initialization vector used to introduce randomness in the encryption process, enhancing security.
func (rcv *EME) IV() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Initialization vector used to introduce randomness in the encryption process, enhancing security.
/// Identifier for the public key used, aiding in recipient key management and message decryption.
func (rcv *EME) PUBLIC_KEY_IDENTIFIER() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(16))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Identifier for the public key used, aiding in recipient key management and message decryption.
/// Specifies the set of cryptographic algorithms used in the encryption process.
func (rcv *EME) CIPHER_SUITE() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(18))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Specifies the set of cryptographic algorithms used in the encryption process.
/// Parameters for the Key Derivation Function, guiding the process of deriving keys from the shared secret.
func (rcv *EME) KDF_PARAMETERS() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(20))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Parameters for the Key Derivation Function, guiding the process of deriving keys from the shared secret.
/// Parameters defining specific settings for the encryption algorithm, such as block size or operation mode.
func (rcv *EME) ENCRYPTION_ALGORITHM_PARAMETERS() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(22))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

/// Parameters defining specific settings for the encryption algorithm, such as block size or operation mode.
func EMEStart(builder *flatbuffers.Builder) {
	builder.StartObject(10)
}
func EMEAddENCRYPTED_BLOB(builder *flatbuffers.Builder, ENCRYPTED_BLOB flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(ENCRYPTED_BLOB), 0)
}
func EMEStartENCRYPTED_BLOBVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func EMEAddEPHEMERAL_PUBLIC_KEY(builder *flatbuffers.Builder, EPHEMERAL_PUBLIC_KEY flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(EPHEMERAL_PUBLIC_KEY), 0)
}
func EMEAddMAC(builder *flatbuffers.Builder, MAC flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(MAC), 0)
}
func EMEAddNONCE(builder *flatbuffers.Builder, NONCE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(NONCE), 0)
}
func EMEAddTAG(builder *flatbuffers.Builder, TAG flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(TAG), 0)
}
func EMEAddIV(builder *flatbuffers.Builder, IV flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(5, flatbuffers.UOffsetT(IV), 0)
}
func EMEAddPUBLIC_KEY_IDENTIFIER(builder *flatbuffers.Builder, PUBLIC_KEY_IDENTIFIER flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(6, flatbuffers.UOffsetT(PUBLIC_KEY_IDENTIFIER), 0)
}
func EMEAddCIPHER_SUITE(builder *flatbuffers.Builder, CIPHER_SUITE flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(7, flatbuffers.UOffsetT(CIPHER_SUITE), 0)
}
func EMEAddKDF_PARAMETERS(builder *flatbuffers.Builder, KDF_PARAMETERS flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(8, flatbuffers.UOffsetT(KDF_PARAMETERS), 0)
}
func EMEAddENCRYPTION_ALGORITHM_PARAMETERS(builder *flatbuffers.Builder, ENCRYPTION_ALGORITHM_PARAMETERS flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(9, flatbuffers.UOffsetT(ENCRYPTION_ALGORITHM_PARAMETERS), 0)
}
func EMEEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
