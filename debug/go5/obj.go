// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs -- -I/opt/remy/go/include -I/opt/remy/go/src/cmd/5l go5/obj.in.go

package go5

const (
	AXXX                = 0x0
	AAND                = 0x1
	AEOR                = 0x2
	ASUB                = 0x3
	ARSB                = 0x4
	AADD                = 0x5
	AADC                = 0x6
	ASBC                = 0x7
	ARSC                = 0x8
	ATST                = 0x9
	ATEQ                = 0xa
	ACMP                = 0xb
	ACMN                = 0xc
	AORR                = 0xd
	ABIC                = 0xe
	AMVN                = 0xf
	AB                  = 0x10
	ABL                 = 0x11
	ABEQ                = 0x12
	ABNE                = 0x13
	ABCS                = 0x14
	ABHS                = 0x15
	ABCC                = 0x16
	ABLO                = 0x17
	ABMI                = 0x18
	ABPL                = 0x19
	ABVS                = 0x1a
	ABVC                = 0x1b
	ABHI                = 0x1c
	ABLS                = 0x1d
	ABGE                = 0x1e
	ABLT                = 0x1f
	ABGT                = 0x20
	ABLE                = 0x21
	AMOVWD              = 0x22
	AMOVWF              = 0x23
	AMOVDW              = 0x24
	AMOVFW              = 0x25
	AMOVFD              = 0x26
	AMOVDF              = 0x27
	AMOVF               = 0x28
	AMOVD               = 0x29
	ACMPF               = 0x2a
	ACMPD               = 0x2b
	AADDF               = 0x2c
	AADDD               = 0x2d
	ASUBF               = 0x2e
	ASUBD               = 0x2f
	AMULF               = 0x30
	AMULD               = 0x31
	ADIVF               = 0x32
	ADIVD               = 0x33
	ASQRTF              = 0x34
	ASQRTD              = 0x35
	AABSF               = 0x36
	AABSD               = 0x37
	ASRL                = 0x38
	ASRA                = 0x39
	ASLL                = 0x3a
	AMULU               = 0x3b
	ADIVU               = 0x3c
	AMUL                = 0x3d
	ADIV                = 0x3e
	AMOD                = 0x3f
	AMODU               = 0x40
	AMOVB               = 0x41
	AMOVBU              = 0x42
	AMOVH               = 0x43
	AMOVHU              = 0x44
	AMOVW               = 0x45
	AMOVM               = 0x46
	ASWPBU              = 0x47
	ASWPW               = 0x48
	ANOP                = 0x49
	ARFE                = 0x4a
	ASWI                = 0x4b
	AMULA               = 0x4c
	ADATA               = 0x4d
	AGLOBL              = 0x4e
	AGOK                = 0x4f
	AHISTORY            = 0x50
	ANAME               = 0x51
	ARET                = 0x52
	ATEXT               = 0x53
	AWORD               = 0x54
	ADYNT_              = 0x55
	AINIT_              = 0x56
	ABCASE              = 0x57
	ACASE               = 0x58
	AEND                = 0x59
	AMULL               = 0x5a
	AMULAL              = 0x5b
	AMULLU              = 0x5c
	AMULALU             = 0x5d
	ABX                 = 0x5e
	ABXRET              = 0x5f
	ADWORD              = 0x60
	ASIGNAME            = 0x61
	ALDREX              = 0x62
	ASTREX              = 0x63
	ALDREXD             = 0x64
	ASTREXD             = 0x65
	APLD                = 0x66
	AUNDEF              = 0x67
	ACLZ                = 0x68
	AMULWT              = 0x69
	AMULWB              = 0x6a
	AMULAWT             = 0x6b
	AMULAWB             = 0x6c
	AUSEFIELD           = 0x6d
	ALAST               = 0x6e
	C_SCOND      Suffix = 0xf
	C_SBIT       Suffix = 0x10
	C_PBIT       Suffix = 0x20
	C_WBIT       Suffix = 0x40
	C_FBIT       Suffix = 0x80
	C_UBIT       Suffix = 0x80
	C_SCOND_EQ   Suffix = 0x0
	C_SCOND_NE   Suffix = 0x1
	C_SCOND_HS   Suffix = 0x2
	C_SCOND_LO   Suffix = 0x3
	C_SCOND_MI   Suffix = 0x4
	C_SCOND_PL   Suffix = 0x5
	C_SCOND_VS   Suffix = 0x6
	C_SCOND_VC   Suffix = 0x7
	C_SCOND_HI   Suffix = 0x8
	C_SCOND_LS   Suffix = 0x9
	C_SCOND_GE   Suffix = 0xa
	C_SCOND_LT   Suffix = 0xb
	C_SCOND_GT   Suffix = 0xc
	C_SCOND_LE   Suffix = 0xd
	C_SCOND_NONE Suffix = 0xe
	C_SCOND_NV   Suffix = 0xf
	D_GOK               = 0x0
	D_NONE              = 0x1
	D_BRANCH            = 0x2
	D_OREG              = 0x3
	D_CONST             = 0x8
	D_FCONST            = 0x9
	D_SCONST            = 0xa
	D_PSR               = 0xb
	D_REG               = 0xd
	D_FREG              = 0xe
	D_FILE              = 0x11
	D_OCONST            = 0x12
	D_FILE1             = 0x13
	D_SHIFT             = 0x14
	D_FPCR              = 0x15
	D_REGREG            = 0x16
	D_ADDR              = 0x17
	D_SBIG              = 0x18
	D_CONST2            = 0x19
	D_REGREG2           = 0x1a
	D_EXTERN            = 0x4
	D_STATIC            = 0x5
	D_AUTO              = 0x6
	D_PARAM             = 0x7
	D_SIZE              = 0x29
	D_PCREL             = 0x2a
	D_GOTOFF            = 0x2b
	D_PLT0              = 0x2c
	D_PLT1              = 0x2d
	D_PLT2              = 0x2e
	D_CALL              = 0x2f
)
