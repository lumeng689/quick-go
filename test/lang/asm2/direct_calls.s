"".Add STEXT nosplit size=15 args=0x10 locals=0x0 funcid=0x0
	0x0000 00000 (direct_calls.go:4)	TEXT	"".Add(SB), NOSPLIT|ABIInternal, $0-16
	0x0000 00000 (direct_calls.go:4)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_calls.go:4)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_calls.go:4)	MOVL	"".b+12(SP), AX
	0x0004 00004 (direct_calls.go:4)	MOVL	"".a+8(SP), CX
	0x0008 00008 (direct_calls.go:4)	ADDL	CX, AX
	0x000a 00010 (direct_calls.go:4)	MOVL	AX, "".~r2+16(SP)
	0x000e 00014 (direct_calls.go:4)	RET
	0x0000 8b 44 24 0c 8b 4c 24 08 01 c8 89 44 24 10 c3     .D$..L$....D$..
"".(*Adder).AddPtr STEXT nosplit size=15 args=0x18 locals=0x0 funcid=0x0
	0x0000 00000 (direct_calls.go:8)	TEXT	"".(*Adder).AddPtr(SB), NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (direct_calls.go:8)	FUNCDATA	$0, gclocals·2a5305abe05176240e61b8620e19a815(SB)
	0x0000 00000 (direct_calls.go:8)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_calls.go:8)	MOVL	"".b+20(SP), AX
	0x0004 00004 (direct_calls.go:8)	MOVL	"".a+16(SP), CX
	0x0008 00008 (direct_calls.go:8)	ADDL	CX, AX
	0x000a 00010 (direct_calls.go:8)	MOVL	AX, "".~r2+24(SP)
	0x000e 00014 (direct_calls.go:8)	RET
	0x0000 8b 44 24 14 8b 4c 24 10 01 c8 89 44 24 18 c3     .D$..L$....D$..
"".Adder.AddVal STEXT nosplit size=15 args=0x18 locals=0x0 funcid=0x0
	0x0000 00000 (direct_calls.go:10)	TEXT	"".Adder.AddVal(SB), NOSPLIT|ABIInternal, $0-24
	0x0000 00000 (direct_calls.go:10)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_calls.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (direct_calls.go:10)	MOVL	"".b+16(SP), AX
	0x0004 00004 (direct_calls.go:10)	MOVL	"".a+12(SP), CX
	0x0008 00008 (direct_calls.go:10)	ADDL	CX, AX
	0x000a 00010 (direct_calls.go:10)	MOVL	AX, "".~r2+24(SP)
	0x000e 00014 (direct_calls.go:10)	RET
	0x0000 8b 44 24 10 8b 4c 24 0c 01 c8 89 44 24 18 c3     .D$..L$....D$..
"".main STEXT size=185 args=0x0 locals=0x28 funcid=0x0
	0x0000 00000 (direct_calls.go:15)	TEXT	"".main(SB), ABIInternal, $40-0
	0x0000 00000 (direct_calls.go:15)	MOVQ	(TLS), CX
	0x0009 00009 (direct_calls.go:15)	CMPQ	SP, 16(CX)
	0x000d 00013 (direct_calls.go:15)	PCDATA	$0, $-2
	0x000d 00013 (direct_calls.go:15)	JLS	175
	0x0013 00019 (direct_calls.go:15)	PCDATA	$0, $-1
	0x0013 00019 (direct_calls.go:15)	SUBQ	$40, SP
	0x0017 00023 (direct_calls.go:15)	MOVQ	BP, 32(SP)
	0x001c 00028 (direct_calls.go:15)	LEAQ	32(SP), BP
	0x0021 00033 (direct_calls.go:15)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0021 00033 (direct_calls.go:15)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0021 00033 (direct_calls.go:16)	MOVQ	$137438953482, AX
	0x002b 00043 (direct_calls.go:16)	MOVQ	AX, (SP)
	0x002f 00047 (direct_calls.go:16)	PCDATA	$1, $0
	0x002f 00047 (direct_calls.go:16)	CALL	"".Add(SB)
	0x0034 00052 (direct_calls.go:18)	MOVL	$0, "".adder+28(SP)
	0x003c 00060 (direct_calls.go:18)	MOVL	$6754, "".adder+28(SP)
	0x0044 00068 (direct_calls.go:19)	LEAQ	"".adder+28(SP), AX
	0x0049 00073 (direct_calls.go:19)	MOVQ	AX, (SP)
	0x004d 00077 (direct_calls.go:19)	MOVQ	$137438953482, AX
	0x0057 00087 (direct_calls.go:19)	MOVQ	AX, 8(SP)
	0x005c 00092 (direct_calls.go:19)	NOP
	0x0060 00096 (direct_calls.go:19)	CALL	"".(*Adder).AddPtr(SB)
	0x0065 00101 (direct_calls.go:20)	MOVL	"".adder+28(SP), AX
	0x0069 00105 (direct_calls.go:20)	MOVL	AX, (SP)
	0x006c 00108 (direct_calls.go:20)	MOVQ	$137438953482, AX
	0x0076 00118 (direct_calls.go:20)	MOVQ	AX, 4(SP)
	0x007b 00123 (direct_calls.go:20)	NOP
	0x0080 00128 (direct_calls.go:20)	CALL	"".Adder.AddVal(SB)
	0x0085 00133 (direct_calls.go:22)	MOVL	"".adder+28(SP), AX
	0x0089 00137 (direct_calls.go:22)	MOVL	AX, (SP)
	0x008c 00140 (direct_calls.go:22)	MOVQ	$137438953482, AX
	0x0096 00150 (direct_calls.go:22)	MOVQ	AX, 4(SP)
	0x009b 00155 (direct_calls.go:22)	NOP
	0x00a0 00160 (direct_calls.go:22)	CALL	"".Adder.AddVal(SB)
	0x00a5 00165 (direct_calls.go:23)	MOVQ	32(SP), BP
	0x00aa 00170 (direct_calls.go:23)	ADDQ	$40, SP
	0x00ae 00174 (direct_calls.go:23)	RET
	0x00af 00175 (direct_calls.go:23)	NOP
	0x00af 00175 (direct_calls.go:15)	PCDATA	$1, $-1
	0x00af 00175 (direct_calls.go:15)	PCDATA	$0, $-2
	0x00af 00175 (direct_calls.go:15)	CALL	runtime.morestack_noctxt(SB)
	0x00b4 00180 (direct_calls.go:15)	PCDATA	$0, $-1
	0x00b4 00180 (direct_calls.go:15)	JMP	0
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 0f 86 9c  dH..%....H;a....
	0x0010 00 00 00 48 83 ec 28 48 89 6c 24 20 48 8d 6c 24  ...H..(H.l$ H.l$
	0x0020 20 48 b8 0a 00 00 00 20 00 00 00 48 89 04 24 e8   H..... ...H..$.
	0x0030 00 00 00 00 c7 44 24 1c 00 00 00 00 c7 44 24 1c  .....D$......D$.
	0x0040 62 1a 00 00 48 8d 44 24 1c 48 89 04 24 48 b8 0a  b...H.D$.H..$H..
	0x0050 00 00 00 20 00 00 00 48 89 44 24 08 0f 1f 40 00  ... ...H.D$...@.
	0x0060 e8 00 00 00 00 8b 44 24 1c 89 04 24 48 b8 0a 00  ......D$...$H...
	0x0070 00 00 20 00 00 00 48 89 44 24 04 0f 1f 44 00 00  .. ...H.D$...D..
	0x0080 e8 00 00 00 00 8b 44 24 1c 89 04 24 48 b8 0a 00  ......D$...$H...
	0x0090 00 00 20 00 00 00 48 89 44 24 04 0f 1f 44 00 00  .. ...H.D$...D..
	0x00a0 e8 00 00 00 00 48 8b 6c 24 20 48 83 c4 28 c3 e8  .....H.l$ H..(..
	0x00b0 00 00 00 00 e9 47 ff ff ff                       .....G...
	rel 5+4 t=17 TLS+0
	rel 48+4 t=8 "".Add+0
	rel 97+4 t=8 "".(*Adder).AddPtr+0
	rel 129+4 t=8 "".Adder.AddVal+0
	rel 161+4 t=8 "".Adder.AddVal+0
	rel 176+4 t=8 runtime.morestack_noctxt+0
"".(*Adder).AddVal STEXT dupok size=124 args=0x18 locals=0x20 funcid=0x16
	0x0000 00000 (<autogenerated>:1)	TEXT	"".(*Adder).AddVal(SB), DUPOK|WRAPPER|ABIInternal, $32-24
	0x0000 00000 (<autogenerated>:1)	MOVQ	(TLS), CX
	0x0009 00009 (<autogenerated>:1)	CMPQ	SP, 16(CX)
	0x000d 00013 (<autogenerated>:1)	PCDATA	$0, $-2
	0x000d 00013 (<autogenerated>:1)	JLS	102
	0x000f 00015 (<autogenerated>:1)	PCDATA	$0, $-1
	0x000f 00015 (<autogenerated>:1)	SUBQ	$32, SP
	0x0013 00019 (<autogenerated>:1)	MOVQ	BP, 24(SP)
	0x0018 00024 (<autogenerated>:1)	LEAQ	24(SP), BP
	0x001d 00029 (<autogenerated>:1)	MOVQ	32(CX), BX
	0x0021 00033 (<autogenerated>:1)	TESTQ	BX, BX
	0x0024 00036 (<autogenerated>:1)	JNE	109
	0x0026 00038 (<autogenerated>:1)	NOP
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$0, gclocals·1a65e721a2ccc325b382662e7ffee780(SB)
	0x0026 00038 (<autogenerated>:1)	FUNCDATA	$1, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0026 00038 (<autogenerated>:1)	MOVQ	""..this+40(SP), AX
	0x002b 00043 (<autogenerated>:1)	TESTQ	AX, AX
	0x002e 00046 (<autogenerated>:1)	JEQ	96
	0x0030 00048 (<autogenerated>:1)	MOVL	(AX), AX
	0x0032 00050 (<autogenerated>:1)	MOVL	AX, (SP)
	0x0035 00053 (<autogenerated>:1)	MOVL	"".a+48(SP), AX
	0x0039 00057 (<autogenerated>:1)	MOVL	AX, 4(SP)
	0x003d 00061 (<autogenerated>:1)	MOVL	"".b+52(SP), AX
	0x0041 00065 (<autogenerated>:1)	MOVL	AX, 8(SP)
	0x0045 00069 (<autogenerated>:1)	PCDATA	$1, $1
	0x0045 00069 (<autogenerated>:1)	CALL	"".Adder.AddVal(SB)
	0x004a 00074 (<autogenerated>:1)	MOVL	16(SP), AX
	0x004e 00078 (<autogenerated>:1)	MOVL	AX, "".~r2+56(SP)
	0x0052 00082 (<autogenerated>:1)	MOVQ	24(SP), BP
	0x0057 00087 (<autogenerated>:1)	ADDQ	$32, SP
	0x005b 00091 (<autogenerated>:1)	RET
	0x005c 00092 (<autogenerated>:1)	NOP
	0x0060 00096 (<autogenerated>:1)	CALL	runtime.panicwrap(SB)
	0x0065 00101 (<autogenerated>:1)	XCHGL	AX, AX
	0x0066 00102 (<autogenerated>:1)	NOP
	0x0066 00102 (<autogenerated>:1)	PCDATA	$1, $-1
	0x0066 00102 (<autogenerated>:1)	PCDATA	$0, $-2
	0x0066 00102 (<autogenerated>:1)	CALL	runtime.morestack_noctxt(SB)
	0x006b 00107 (<autogenerated>:1)	PCDATA	$0, $-1
	0x006b 00107 (<autogenerated>:1)	JMP	0
	0x006d 00109 (<autogenerated>:1)	LEAQ	40(SP), DI
	0x0072 00114 (<autogenerated>:1)	CMPQ	(BX), DI
	0x0075 00117 (<autogenerated>:1)	JNE	38
	0x0077 00119 (<autogenerated>:1)	MOVQ	SP, (BX)
	0x007a 00122 (<autogenerated>:1)	JMP	38
	0x0000 64 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 57 48  dH..%....H;a.vWH
	0x0010 83 ec 20 48 89 6c 24 18 48 8d 6c 24 18 48 8b 59  .. H.l$.H.l$.H.Y
	0x0020 20 48 85 db 75 47 48 8b 44 24 28 48 85 c0 74 30   H..uGH.D$(H..t0
	0x0030 8b 00 89 04 24 8b 44 24 30 89 44 24 04 8b 44 24  ....$.D$0.D$..D$
	0x0040 34 89 44 24 08 e8 00 00 00 00 8b 44 24 10 89 44  4.D$.......D$..D
	0x0050 24 38 48 8b 6c 24 18 48 83 c4 20 c3 0f 1f 40 00  $8H.l$.H.. ...@.
	0x0060 e8 00 00 00 00 90 e8 00 00 00 00 eb 93 48 8d 7c  .............H.|
	0x0070 24 28 48 39 3b 75 af 48 89 23 eb aa              $(H9;u.H.#..
	rel 5+4 t=17 TLS+0
	rel 70+4 t=8 "".Adder.AddVal+0
	rel 97+4 t=8 runtime.panicwrap+0
	rel 103+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
	0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 00 00 00 00 00 00 00 00                          ........
runtime.memequal32·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal32+0
runtime.memequal64·f SRODATA dupok size=8
	0x0000 00 00 00 00 00 00 00 00                          ........
	rel 0+8 t=1 runtime.memequal64+0
runtime.gcbits.01 SRODATA dupok size=1
	0x0000 01                                               .
type..namedata.*main.Adder. SRODATA dupok size=14
	0x0000 01 00 0b 2a 6d 61 69 6e 2e 41 64 64 65 72        ...*main.Adder
type..namedata.*func(*main.Adder, int32, int32) int32- SRODATA dupok size=41
	0x0000 00 00 26 2a 66 75 6e 63 28 2a 6d 61 69 6e 2e 41  ..&*func(*main.A
	0x0010 64 64 65 72 2c 20 69 6e 74 33 32 2c 20 69 6e 74  dder, int32, int
	0x0020 33 32 29 20 69 6e 74 33 32                       32) int32
type.*func(*"".Adder, int32, int32) int32 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 35 26 2b 33 08 08 08 36 00 00 00 00 00 00 00 00  5&+3...6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Adder, int32, int32) int32-+0
	rel 48+8 t=1 type.func(*"".Adder, int32, int32) int32+0
type.func(*"".Adder, int32, int32) int32 SRODATA dupok size=88
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 d7 38 ec a1 02 08 08 33 00 00 00 00 00 00 00 00  .8.....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 03 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(*main.Adder, int32, int32) int32-+0
	rel 44+4 t=6 type.*func(*"".Adder, int32, int32) int32+0
	rel 56+8 t=1 type.*"".Adder+0
	rel 64+8 t=1 type.int32+0
	rel 72+8 t=1 type.int32+0
	rel 80+8 t=1 type.int32+0
type..namedata.AddPtr. SRODATA dupok size=9
	0x0000 01 00 06 41 64 64 50 74 72                       ...AddPtr
type..namedata.*func(int32, int32) int32- SRODATA dupok size=28
	0x0000 00 00 19 2a 66 75 6e 63 28 69 6e 74 33 32 2c 20  ...*func(int32, 
	0x0010 69 6e 74 33 32 29 20 69 6e 74 33 32              int32) int32
type.*func(int32, int32) int32 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 c0 41 58 87 08 08 08 36 00 00 00 00 00 00 00 00  .AX....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int32, int32) int32-+0
	rel 48+8 t=1 type.func(int32, int32) int32+0
type.func(int32, int32) int32 SRODATA dupok size=80
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 13 9f 60 7f 02 08 08 33 00 00 00 00 00 00 00 00  ..`....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 02 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(int32, int32) int32-+0
	rel 44+4 t=6 type.*func(int32, int32) int32+0
	rel 56+8 t=1 type.int32+0
	rel 64+8 t=1 type.int32+0
	rel 72+8 t=1 type.int32+0
type..namedata.AddVal. SRODATA dupok size=9
	0x0000 01 00 06 41 64 64 56 61 6c                       ...AddVal
type.*"".Adder SRODATA size=104
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 46 c2 6b c4 09 08 08 36 00 00 00 00 00 00 00 00  F.k....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 02 00 02 00  ................
	0x0040 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0060 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*main.Adder.+0
	rel 48+8 t=1 type."".Adder+0
	rel 56+4 t=5 type..importpath."".+0
	rel 72+4 t=5 type..namedata.AddPtr.+0
	rel 76+4 t=27 type.func(int32, int32) int32+0
	rel 80+4 t=27 "".(*Adder).AddPtr+0
	rel 84+4 t=27 "".(*Adder).AddPtr+0
	rel 88+4 t=5 type..namedata.AddVal.+0
	rel 92+4 t=27 type.func(int32, int32) int32+0
	rel 96+4 t=27 "".(*Adder).AddVal+0
	rel 100+4 t=27 "".(*Adder).AddVal+0
runtime.gcbits. SRODATA dupok size=0
type..namedata.*func(main.Adder, int32, int32) int32- SRODATA dupok size=40
	0x0000 00 00 25 2a 66 75 6e 63 28 6d 61 69 6e 2e 41 64  ..%*func(main.Ad
	0x0010 64 65 72 2c 20 69 6e 74 33 32 2c 20 69 6e 74 33  der, int32, int3
	0x0020 32 29 20 69 6e 74 33 32                          2) int32
type.*func("".Adder, int32, int32) int32 SRODATA dupok size=56
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 ec 3b d4 ed 08 08 08 36 00 00 00 00 00 00 00 00  .;.....6........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal64·f+0
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.Adder, int32, int32) int32-+0
	rel 48+8 t=1 type.func("".Adder, int32, int32) int32+0
type.func("".Adder, int32, int32) int32 SRODATA dupok size=88
	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
	0x0010 80 6d 70 15 02 08 08 33 00 00 00 00 00 00 00 00  .mp....3........
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 03 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 00 00 00 00                          ........
	rel 32+8 t=1 runtime.gcbits.01+0
	rel 40+4 t=5 type..namedata.*func(main.Adder, int32, int32) int32-+0
	rel 44+4 t=6 type.*func("".Adder, int32, int32) int32+0
	rel 56+8 t=1 type."".Adder+0
	rel 64+8 t=1 type.int32+0
	rel 72+8 t=1 type.int32+0
	rel 80+8 t=1 type.int32+0
type..namedata.id- SRODATA dupok size=5
	0x0000 00 00 02 69 64                                   ...id
type."".Adder SRODATA size=136
	0x0000 04 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0010 8a 3d 5f 61 0f 04 04 19 00 00 00 00 00 00 00 00  .=_a............
	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0040 01 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
	0x0050 00 00 00 00 01 00 01 00 28 00 00 00 00 00 00 00  ........(.......
	0x0060 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0070 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	0x0080 00 00 00 00 00 00 00 00                          ........
	rel 24+8 t=1 runtime.memequal32·f+0
	rel 32+8 t=1 runtime.gcbits.+0
	rel 40+4 t=5 type..namedata.*main.Adder.+0
	rel 44+4 t=5 type.*"".Adder+0
	rel 48+8 t=1 type..importpath."".+0
	rel 56+8 t=1 type."".Adder+96
	rel 80+4 t=5 type..importpath."".+0
	rel 96+8 t=1 type..namedata.id-+0
	rel 104+8 t=1 type.int32+0
	rel 120+4 t=5 type..namedata.AddVal.+0
	rel 124+4 t=27 type.func(int32, int32) int32+0
	rel 128+4 t=27 "".(*Adder).AddVal+0
	rel 132+4 t=27 "".Adder.AddVal+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
	0x0000 01 00 00 00 00 00 00 00                          ........
gclocals·2a5305abe05176240e61b8620e19a815 SRODATA dupok size=9
	0x0000 01 00 00 00 01 00 00 00 00                       .........
gclocals·1a65e721a2ccc325b382662e7ffee780 SRODATA dupok size=10
	0x0000 02 00 00 00 01 00 00 00 01 00                    ..........
gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	0x0000 02 00 00 00 00 00 00 00                          ........
