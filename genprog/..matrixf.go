package mathgl

import(
	 //"math"
)

type Mat2f [4]float32
type Mat2x3f [6]float32
type Mat2x4f [8]float32
type Mat3x2f [6]float32
type Mat3f [9]float32
type Mat3x4f [12]float32
type Mat4x2f [8]float32
type Mat4x3f [12]float32
type Mat4f [16]float32

func Ident2f() Mat2f {
	return Mat2f{1,0,0,1}
}

func Ident3f() Mat3f {
	return Mat3f{1,0,0,0,1,0,0,0,1}
}

func Ident4f() Mat4f {
	return Mat4f{1,0,0,0,0,1,0,0,0,0,1,0,0,0,0,1}
}

func (m1 Mat2f) Add(m2 Mat2f) Mat2f {
	return Mat2f {m1[0] + m2[0],m1[1] + m2[1],m1[2] + m2[2],m1[3] + m2[3]}
}

func (m1 Mat2x3f) Add(m2 Mat2x3f) Mat2x3f {
	return Mat2x3f {m1[0] + m2[0],m1[1] + m2[1],m1[2] + m2[2],m1[3] + m2[3],m1[4] + m2[4],m1[5] + m2[5]}
}

func (m1 Mat2x4f) Add(m2 Mat2x4f) Mat2x4f {
	return Mat2x4f {m1[0] + m2[0],m1[1] + m2[1],m1[2] + m2[2],m1[3] + m2[3],m1[4] + m2[4],m1[5] + m2[5],m1[6] + m2[6],m1[7] + m2[7]}
}

func (m1 Mat3x2f) Add(m2 Mat3x2f) Mat3x2f {
	return Mat3x2f {m1[0] + m2[0],m1[1] + m2[1],m1[2] + m2[2],m1[3] + m2[3],m1[4] + m2[4],m1[5] + m2[5]}
}

func (m1 Mat3f) Add(m2 Mat3f) Mat3f {
	return Mat3f {m1[0] + m2[0],m1[1] + m2[1],m1[2] + m2[2],m1[3] + m2[3],m1[4] + m2[4],m1[5] + m2[5],m1[6] + m2[6],m1[7] + m2[7],m1[8] + m2[8]}
}

func (m1 Mat3x4f) Add(m2 Mat3x4f) Mat3x4f {
	return Mat3x4f {m1[0] + m2[0],m1[1] + m2[1],m1[2] + m2[2],m1[3] + m2[3],m1[4] + m2[4],m1[5] + m2[5],m1[6] + m2[6],m1[7] + m2[7],m1[8] + m2[8],m1[9] + m2[9],m1[10] + m2[10],m1[11] + m2[11]}
}

func (m1 Mat4x2f) Add(m2 Mat4x2f) Mat4x2f {
	return Mat4x2f {m1[0] + m2[0],m1[1] + m2[1],m1[2] + m2[2],m1[3] + m2[3],m1[4] + m2[4],m1[5] + m2[5],m1[6] + m2[6],m1[7] + m2[7]}
}

func (m1 Mat4x3f) Add(m2 Mat4x3f) Mat4x3f {
	return Mat4x3f {m1[0] + m2[0],m1[1] + m2[1],m1[2] + m2[2],m1[3] + m2[3],m1[4] + m2[4],m1[5] + m2[5],m1[6] + m2[6],m1[7] + m2[7],m1[8] + m2[8],m1[9] + m2[9],m1[10] + m2[10],m1[11] + m2[11]}
}

func (m1 Mat4f) Add(m2 Mat4f) Mat4f {
	return Mat4f {m1[0] + m2[0],m1[1] + m2[1],m1[2] + m2[2],m1[3] + m2[3],m1[4] + m2[4],m1[5] + m2[5],m1[6] + m2[6],m1[7] + m2[7],m1[8] + m2[8],m1[9] + m2[9],m1[10] + m2[10],m1[11] + m2[11],m1[12] + m2[12],m1[13] + m2[13],m1[14] + m2[14],m1[15] + m2[15]}
}

func (m1 Mat2f) Sub(m2 Mat2f) Mat2f {
	return Mat2f {m1[0] - m2[0],m1[1] - m2[1],m1[2] - m2[2],m1[3] - m2[3]}
}

func (m1 Mat2x3f) Sub(m2 Mat2x3f) Mat2x3f {
	return Mat2x3f {m1[0] - m2[0],m1[1] - m2[1],m1[2] - m2[2],m1[3] - m2[3],m1[4] - m2[4],m1[5] - m2[5]}
}

func (m1 Mat2x4f) Sub(m2 Mat2x4f) Mat2x4f {
	return Mat2x4f {m1[0] - m2[0],m1[1] - m2[1],m1[2] - m2[2],m1[3] - m2[3],m1[4] - m2[4],m1[5] - m2[5],m1[6] - m2[6],m1[7] - m2[7]}
}

func (m1 Mat3x2f) Sub(m2 Mat3x2f) Mat3x2f {
	return Mat3x2f {m1[0] - m2[0],m1[1] - m2[1],m1[2] - m2[2],m1[3] - m2[3],m1[4] - m2[4],m1[5] - m2[5]}
}

func (m1 Mat3f) Sub(m2 Mat3f) Mat3f {
	return Mat3f {m1[0] - m2[0],m1[1] - m2[1],m1[2] - m2[2],m1[3] - m2[3],m1[4] - m2[4],m1[5] - m2[5],m1[6] - m2[6],m1[7] - m2[7],m1[8] - m2[8]}
}

func (m1 Mat3x4f) Sub(m2 Mat3x4f) Mat3x4f {
	return Mat3x4f {m1[0] - m2[0],m1[1] - m2[1],m1[2] - m2[2],m1[3] - m2[3],m1[4] - m2[4],m1[5] - m2[5],m1[6] - m2[6],m1[7] - m2[7],m1[8] - m2[8],m1[9] - m2[9],m1[10] - m2[10],m1[11] - m2[11]}
}

func (m1 Mat4x2f) Sub(m2 Mat4x2f) Mat4x2f {
	return Mat4x2f {m1[0] - m2[0],m1[1] - m2[1],m1[2] - m2[2],m1[3] - m2[3],m1[4] - m2[4],m1[5] - m2[5],m1[6] - m2[6],m1[7] - m2[7]}
}

func (m1 Mat4x3f) Sub(m2 Mat4x3f) Mat4x3f {
	return Mat4x3f {m1[0] - m2[0],m1[1] - m2[1],m1[2] - m2[2],m1[3] - m2[3],m1[4] - m2[4],m1[5] - m2[5],m1[6] - m2[6],m1[7] - m2[7],m1[8] - m2[8],m1[9] - m2[9],m1[10] - m2[10],m1[11] - m2[11]}
}

func (m1 Mat4f) Sub(m2 Mat4f) Mat4f {
	return Mat4f {m1[0] - m2[0],m1[1] - m2[1],m1[2] - m2[2],m1[3] - m2[3],m1[4] - m2[4],m1[5] - m2[5],m1[6] - m2[6],m1[7] - m2[7],m1[8] - m2[8],m1[9] - m2[9],m1[10] - m2[10],m1[11] - m2[11],m1[12] - m2[12],m1[13] - m2[13],m1[14] - m2[14],m1[15] - m2[15]}
}

func (m1 Mat2f) Mul(c float32) Mat2f {
	return Mat2f{m1[0] *c,m1[1] *c,m1[2] *c,m1[3] *c}
}

func (m1 Mat2x3f) Mul(c float32) Mat2x3f {
	return Mat2x3f{m1[0] *c,m1[1] *c,m1[2] *c,m1[3] *c,m1[4] *c,m1[5] *c}
}

func (m1 Mat2x4f) Mul(c float32) Mat2x4f {
	return Mat2x4f{m1[0] *c,m1[1] *c,m1[2] *c,m1[3] *c,m1[4] *c,m1[5] *c,m1[6] *c,m1[7] *c}
}

func (m1 Mat3x2f) Mul(c float32) Mat3x2f {
	return Mat3x2f{m1[0] *c,m1[1] *c,m1[2] *c,m1[3] *c,m1[4] *c,m1[5] *c}
}

func (m1 Mat3f) Mul(c float32) Mat3f {
	return Mat3f{m1[0] *c,m1[1] *c,m1[2] *c,m1[3] *c,m1[4] *c,m1[5] *c,m1[6] *c,m1[7] *c,m1[8] *c}
}

func (m1 Mat3x4f) Mul(c float32) Mat3x4f {
	return Mat3x4f{m1[0] *c,m1[1] *c,m1[2] *c,m1[3] *c,m1[4] *c,m1[5] *c,m1[6] *c,m1[7] *c,m1[8] *c,m1[9] *c,m1[10] *c,m1[11] *c}
}

func (m1 Mat4x2f) Mul(c float32) Mat4x2f {
	return Mat4x2f{m1[0] *c,m1[1] *c,m1[2] *c,m1[3] *c,m1[4] *c,m1[5] *c,m1[6] *c,m1[7] *c}
}

func (m1 Mat4x3f) Mul(c float32) Mat4x3f {
	return Mat4x3f{m1[0] *c,m1[1] *c,m1[2] *c,m1[3] *c,m1[4] *c,m1[5] *c,m1[6] *c,m1[7] *c,m1[8] *c,m1[9] *c,m1[10] *c,m1[11] *c}
}

func (m1 Mat4f) Mul(c float32) Mat4f {
	return Mat4f{m1[0] *c,m1[1] *c,m1[2] *c,m1[3] *c,m1[4] *c,m1[5] *c,m1[6] *c,m1[7] *c,m1[8] *c,m1[9] *c,m1[10] *c,m1[11] *c,m1[12] *c,m1[13] *c,m1[14] *c,m1[15] *c}
}

func (m1 Mat2f) Mul2x1(m2 Vec2f) Vec2f {
	return Vec2f{m1[0] * m2[0] + m1[2] * m2[1], m1[1] * m2[0] + m1[3] * m2[1]}
}

func (m1 Mat2f) Mul2(m2 Mat2f) Mat2f {
	return Mat2f{m1[0] * m2[0] + m1[2] * m2[1], m1[1] * m2[0] + m1[3] * m2[1], m1[0] * m2[2] + m1[2] * m2[3], m1[1] * m2[2] + m1[3] * m2[3]}
}

func (m1 Mat2f) Mul2x3(m2 Mat2x3f) Mat2x3f {
	return Mat2x3f{m1[0] * m2[0] + m1[2] * m2[1], m1[1] * m2[0] + m1[3] * m2[1], m1[0] * m2[2] + m1[2] * m2[3], m1[1] * m2[2] + m1[3] * m2[3], m1[0] * m2[4] + m1[2] * m2[5], m1[1] * m2[4] + m1[3] * m2[5]}
}

func (m1 Mat2f) Mul2x4(m2 Mat2x4f) Mat2x4f {
	return Mat2x4f{m1[0] * m2[0] + m1[2] * m2[1], m1[1] * m2[0] + m1[3] * m2[1], m1[0] * m2[2] + m1[2] * m2[3], m1[1] * m2[2] + m1[3] * m2[3], m1[0] * m2[4] + m1[2] * m2[5], m1[1] * m2[4] + m1[3] * m2[5], m1[0] * m2[6] + m1[2] * m2[7], m1[1] * m2[6] + m1[3] * m2[7]}
}

func (m1 Mat2x3f) Mul3x1(m2 Vec3f) Vec2f {
	return Vec2f{m1[0] * m2[0] + m1[2] * m2[1] + m1[4] * m2[2], m1[1] * m2[0] + m1[3] * m2[1] + m1[5] * m2[2]}
}

func (m1 Mat2x3f) Mul3x2(m2 Mat3x2f) Mat2f {
	return Mat2f{m1[0] * m2[0] + m1[2] * m2[1] + m1[4] * m2[2], m1[1] * m2[0] + m1[3] * m2[1] + m1[5] * m2[2], m1[0] * m2[3] + m1[2] * m2[4] + m1[4] * m2[5], m1[1] * m2[3] + m1[3] * m2[4] + m1[5] * m2[5]}
}

func (m1 Mat2x3f) Mul3(m2 Mat3f) Mat2x3f {
	return Mat2x3f{m1[0] * m2[0] + m1[2] * m2[1] + m1[4] * m2[2], m1[1] * m2[0] + m1[3] * m2[1] + m1[5] * m2[2], m1[0] * m2[3] + m1[2] * m2[4] + m1[4] * m2[5], m1[1] * m2[3] + m1[3] * m2[4] + m1[5] * m2[5], m1[0] * m2[6] + m1[2] * m2[7] + m1[4] * m2[8], m1[1] * m2[6] + m1[3] * m2[7] + m1[5] * m2[8]}
}

func (m1 Mat2x3f) Mul3x4(m2 Mat3x4f) Mat2x4f {
	return Mat2x4f{m1[0] * m2[0] + m1[2] * m2[1] + m1[4] * m2[2], m1[1] * m2[0] + m1[3] * m2[1] + m1[5] * m2[2], m1[0] * m2[3] + m1[2] * m2[4] + m1[4] * m2[5], m1[1] * m2[3] + m1[3] * m2[4] + m1[5] * m2[5], m1[0] * m2[6] + m1[2] * m2[7] + m1[4] * m2[8], m1[1] * m2[6] + m1[3] * m2[7] + m1[5] * m2[8], m1[0] * m2[9] + m1[2] * m2[10] + m1[4] * m2[11], m1[1] * m2[9] + m1[3] * m2[10] + m1[5] * m2[11]}
}

func (m1 Mat2x4f) Mul4x1(m2 Vec4f) Vec2f {
	return Vec2f{m1[0] * m2[0] + m1[2] * m2[1] + m1[4] * m2[2] + m1[6] * m2[3], m1[1] * m2[0] + m1[3] * m2[1] + m1[5] * m2[2] + m1[7] * m2[3]}
}

func (m1 Mat2x4f) Mul4x2(m2 Mat4x2f) Mat2f {
	return Mat2f{m1[0] * m2[0] + m1[2] * m2[1] + m1[4] * m2[2] + m1[6] * m2[3], m1[1] * m2[0] + m1[3] * m2[1] + m1[5] * m2[2] + m1[7] * m2[3], m1[0] * m2[4] + m1[2] * m2[5] + m1[4] * m2[6] + m1[6] * m2[7], m1[1] * m2[4] + m1[3] * m2[5] + m1[5] * m2[6] + m1[7] * m2[7]}
}

func (m1 Mat2x4f) Mul4x3(m2 Mat4x3f) Mat2x3f {
	return Mat2x3f{m1[0] * m2[0] + m1[2] * m2[1] + m1[4] * m2[2] + m1[6] * m2[3], m1[1] * m2[0] + m1[3] * m2[1] + m1[5] * m2[2] + m1[7] * m2[3], m1[0] * m2[4] + m1[2] * m2[5] + m1[4] * m2[6] + m1[6] * m2[7], m1[1] * m2[4] + m1[3] * m2[5] + m1[5] * m2[6] + m1[7] * m2[7], m1[0] * m2[8] + m1[2] * m2[9] + m1[4] * m2[10] + m1[6] * m2[11], m1[1] * m2[8] + m1[3] * m2[9] + m1[5] * m2[10] + m1[7] * m2[11]}
}

func (m1 Mat2x4f) Mul4(m2 Mat4f) Mat2x4f {
	return Mat2x4f{m1[0] * m2[0] + m1[2] * m2[1] + m1[4] * m2[2] + m1[6] * m2[3], m1[1] * m2[0] + m1[3] * m2[1] + m1[5] * m2[2] + m1[7] * m2[3], m1[0] * m2[4] + m1[2] * m2[5] + m1[4] * m2[6] + m1[6] * m2[7], m1[1] * m2[4] + m1[3] * m2[5] + m1[5] * m2[6] + m1[7] * m2[7], m1[0] * m2[8] + m1[2] * m2[9] + m1[4] * m2[10] + m1[6] * m2[11], m1[1] * m2[8] + m1[3] * m2[9] + m1[5] * m2[10] + m1[7] * m2[11], m1[0] * m2[12] + m1[2] * m2[13] + m1[4] * m2[14] + m1[6] * m2[15], m1[1] * m2[12] + m1[3] * m2[13] + m1[5] * m2[14] + m1[7] * m2[15]}
}

func (m1 Mat3x2f) Mul2x1(m2 Vec2f) Vec3f {
	return Vec3f{m1[0] * m2[0] + m1[3] * m2[1], m1[1] * m2[0] + m1[4] * m2[1], m1[2] * m2[0] + m1[5] * m2[1]}
}

func (m1 Mat3x2f) Mul2(m2 Mat2f) Mat3x2f {
	return Mat3x2f{m1[0] * m2[0] + m1[3] * m2[1], m1[1] * m2[0] + m1[4] * m2[1], m1[2] * m2[0] + m1[5] * m2[1], m1[0] * m2[2] + m1[3] * m2[3], m1[1] * m2[2] + m1[4] * m2[3], m1[2] * m2[2] + m1[5] * m2[3]}
}

func (m1 Mat3x2f) Mul2x3(m2 Mat2x3f) Mat3f {
	return Mat3f{m1[0] * m2[0] + m1[3] * m2[1], m1[1] * m2[0] + m1[4] * m2[1], m1[2] * m2[0] + m1[5] * m2[1], m1[0] * m2[2] + m1[3] * m2[3], m1[1] * m2[2] + m1[4] * m2[3], m1[2] * m2[2] + m1[5] * m2[3], m1[0] * m2[4] + m1[3] * m2[5], m1[1] * m2[4] + m1[4] * m2[5], m1[2] * m2[4] + m1[5] * m2[5]}
}

func (m1 Mat3x2f) Mul2x4(m2 Mat2x4f) Mat3x4f {
	return Mat3x4f{m1[0] * m2[0] + m1[3] * m2[1], m1[1] * m2[0] + m1[4] * m2[1], m1[2] * m2[0] + m1[5] * m2[1], m1[0] * m2[2] + m1[3] * m2[3], m1[1] * m2[2] + m1[4] * m2[3], m1[2] * m2[2] + m1[5] * m2[3], m1[0] * m2[4] + m1[3] * m2[5], m1[1] * m2[4] + m1[4] * m2[5], m1[2] * m2[4] + m1[5] * m2[5], m1[0] * m2[6] + m1[3] * m2[7], m1[1] * m2[6] + m1[4] * m2[7], m1[2] * m2[6] + m1[5] * m2[7]}
}

func (m1 Mat3f) Mul3x1(m2 Vec3f) Vec3f {
	return Vec3f{m1[0] * m2[0] + m1[3] * m2[1] + m1[6] * m2[2], m1[1] * m2[0] + m1[4] * m2[1] + m1[7] * m2[2], m1[2] * m2[0] + m1[5] * m2[1] + m1[8] * m2[2]}
}

func (m1 Mat3f) Mul3x2(m2 Mat3x2f) Mat3x2f {
	return Mat3x2f{m1[0] * m2[0] + m1[3] * m2[1] + m1[6] * m2[2], m1[1] * m2[0] + m1[4] * m2[1] + m1[7] * m2[2], m1[2] * m2[0] + m1[5] * m2[1] + m1[8] * m2[2], m1[0] * m2[3] + m1[3] * m2[4] + m1[6] * m2[5], m1[1] * m2[3] + m1[4] * m2[4] + m1[7] * m2[5], m1[2] * m2[3] + m1[5] * m2[4] + m1[8] * m2[5]}
}

func (m1 Mat3f) Mul3(m2 Mat3f) Mat3f {
	return Mat3f{m1[0] * m2[0] + m1[3] * m2[1] + m1[6] * m2[2], m1[1] * m2[0] + m1[4] * m2[1] + m1[7] * m2[2], m1[2] * m2[0] + m1[5] * m2[1] + m1[8] * m2[2], m1[0] * m2[3] + m1[3] * m2[4] + m1[6] * m2[5], m1[1] * m2[3] + m1[4] * m2[4] + m1[7] * m2[5], m1[2] * m2[3] + m1[5] * m2[4] + m1[8] * m2[5], m1[0] * m2[6] + m1[3] * m2[7] + m1[6] * m2[8], m1[1] * m2[6] + m1[4] * m2[7] + m1[7] * m2[8], m1[2] * m2[6] + m1[5] * m2[7] + m1[8] * m2[8]}
}

func (m1 Mat3f) Mul3x4(m2 Mat3x4f) Mat3x4f {
	return Mat3x4f{m1[0] * m2[0] + m1[3] * m2[1] + m1[6] * m2[2], m1[1] * m2[0] + m1[4] * m2[1] + m1[7] * m2[2], m1[2] * m2[0] + m1[5] * m2[1] + m1[8] * m2[2], m1[0] * m2[3] + m1[3] * m2[4] + m1[6] * m2[5], m1[1] * m2[3] + m1[4] * m2[4] + m1[7] * m2[5], m1[2] * m2[3] + m1[5] * m2[4] + m1[8] * m2[5], m1[0] * m2[6] + m1[3] * m2[7] + m1[6] * m2[8], m1[1] * m2[6] + m1[4] * m2[7] + m1[7] * m2[8], m1[2] * m2[6] + m1[5] * m2[7] + m1[8] * m2[8], m1[0] * m2[9] + m1[3] * m2[10] + m1[6] * m2[11], m1[1] * m2[9] + m1[4] * m2[10] + m1[7] * m2[11], m1[2] * m2[9] + m1[5] * m2[10] + m1[8] * m2[11]}
}

func (m1 Mat3x4f) Mul4x1(m2 Vec4f) Vec3f {
	return Vec3f{m1[0] * m2[0] + m1[3] * m2[1] + m1[6] * m2[2] + m1[9] * m2[3], m1[1] * m2[0] + m1[4] * m2[1] + m1[7] * m2[2] + m1[10] * m2[3], m1[2] * m2[0] + m1[5] * m2[1] + m1[8] * m2[2] + m1[11] * m2[3]}
}

func (m1 Mat3x4f) Mul4x2(m2 Mat4x2f) Mat3x2f {
	return Mat3x2f{m1[0] * m2[0] + m1[3] * m2[1] + m1[6] * m2[2] + m1[9] * m2[3], m1[1] * m2[0] + m1[4] * m2[1] + m1[7] * m2[2] + m1[10] * m2[3], m1[2] * m2[0] + m1[5] * m2[1] + m1[8] * m2[2] + m1[11] * m2[3], m1[0] * m2[4] + m1[3] * m2[5] + m1[6] * m2[6] + m1[9] * m2[7], m1[1] * m2[4] + m1[4] * m2[5] + m1[7] * m2[6] + m1[10] * m2[7], m1[2] * m2[4] + m1[5] * m2[5] + m1[8] * m2[6] + m1[11] * m2[7]}
}

func (m1 Mat3x4f) Mul4x3(m2 Mat4x3f) Mat3f {
	return Mat3f{m1[0] * m2[0] + m1[3] * m2[1] + m1[6] * m2[2] + m1[9] * m2[3], m1[1] * m2[0] + m1[4] * m2[1] + m1[7] * m2[2] + m1[10] * m2[3], m1[2] * m2[0] + m1[5] * m2[1] + m1[8] * m2[2] + m1[11] * m2[3], m1[0] * m2[4] + m1[3] * m2[5] + m1[6] * m2[6] + m1[9] * m2[7], m1[1] * m2[4] + m1[4] * m2[5] + m1[7] * m2[6] + m1[10] * m2[7], m1[2] * m2[4] + m1[5] * m2[5] + m1[8] * m2[6] + m1[11] * m2[7], m1[0] * m2[8] + m1[3] * m2[9] + m1[6] * m2[10] + m1[9] * m2[11], m1[1] * m2[8] + m1[4] * m2[9] + m1[7] * m2[10] + m1[10] * m2[11], m1[2] * m2[8] + m1[5] * m2[9] + m1[8] * m2[10] + m1[11] * m2[11]}
}

func (m1 Mat3x4f) Mul4(m2 Mat4f) Mat3x4f {
	return Mat3x4f{m1[0] * m2[0] + m1[3] * m2[1] + m1[6] * m2[2] + m1[9] * m2[3], m1[1] * m2[0] + m1[4] * m2[1] + m1[7] * m2[2] + m1[10] * m2[3], m1[2] * m2[0] + m1[5] * m2[1] + m1[8] * m2[2] + m1[11] * m2[3], m1[0] * m2[4] + m1[3] * m2[5] + m1[6] * m2[6] + m1[9] * m2[7], m1[1] * m2[4] + m1[4] * m2[5] + m1[7] * m2[6] + m1[10] * m2[7], m1[2] * m2[4] + m1[5] * m2[5] + m1[8] * m2[6] + m1[11] * m2[7], m1[0] * m2[8] + m1[3] * m2[9] + m1[6] * m2[10] + m1[9] * m2[11], m1[1] * m2[8] + m1[4] * m2[9] + m1[7] * m2[10] + m1[10] * m2[11], m1[2] * m2[8] + m1[5] * m2[9] + m1[8] * m2[10] + m1[11] * m2[11], m1[0] * m2[12] + m1[3] * m2[13] + m1[6] * m2[14] + m1[9] * m2[15], m1[1] * m2[12] + m1[4] * m2[13] + m1[7] * m2[14] + m1[10] * m2[15], m1[2] * m2[12] + m1[5] * m2[13] + m1[8] * m2[14] + m1[11] * m2[15]}
}

func (m1 Mat4x2f) Mul2x1(m2 Vec2f) Vec4f {
	return Vec4f{m1[0] * m2[0] + m1[4] * m2[1], m1[1] * m2[0] + m1[5] * m2[1], m1[2] * m2[0] + m1[6] * m2[1], m1[3] * m2[0] + m1[7] * m2[1]}
}

func (m1 Mat4x2f) Mul2(m2 Mat2f) Mat4x2f {
	return Mat4x2f{m1[0] * m2[0] + m1[4] * m2[1], m1[1] * m2[0] + m1[5] * m2[1], m1[2] * m2[0] + m1[6] * m2[1], m1[3] * m2[0] + m1[7] * m2[1], m1[0] * m2[2] + m1[4] * m2[3], m1[1] * m2[2] + m1[5] * m2[3], m1[2] * m2[2] + m1[6] * m2[3], m1[3] * m2[2] + m1[7] * m2[3]}
}

func (m1 Mat4x2f) Mul2x3(m2 Mat2x3f) Mat4x3f {
	return Mat4x3f{m1[0] * m2[0] + m1[4] * m2[1], m1[1] * m2[0] + m1[5] * m2[1], m1[2] * m2[0] + m1[6] * m2[1], m1[3] * m2[0] + m1[7] * m2[1], m1[0] * m2[2] + m1[4] * m2[3], m1[1] * m2[2] + m1[5] * m2[3], m1[2] * m2[2] + m1[6] * m2[3], m1[3] * m2[2] + m1[7] * m2[3], m1[0] * m2[4] + m1[4] * m2[5], m1[1] * m2[4] + m1[5] * m2[5], m1[2] * m2[4] + m1[6] * m2[5], m1[3] * m2[4] + m1[7] * m2[5]}
}

func (m1 Mat4x2f) Mul2x4(m2 Mat2x4f) Mat4f {
	return Mat4f{m1[0] * m2[0] + m1[4] * m2[1], m1[1] * m2[0] + m1[5] * m2[1], m1[2] * m2[0] + m1[6] * m2[1], m1[3] * m2[0] + m1[7] * m2[1], m1[0] * m2[2] + m1[4] * m2[3], m1[1] * m2[2] + m1[5] * m2[3], m1[2] * m2[2] + m1[6] * m2[3], m1[3] * m2[2] + m1[7] * m2[3], m1[0] * m2[4] + m1[4] * m2[5], m1[1] * m2[4] + m1[5] * m2[5], m1[2] * m2[4] + m1[6] * m2[5], m1[3] * m2[4] + m1[7] * m2[5], m1[0] * m2[6] + m1[4] * m2[7], m1[1] * m2[6] + m1[5] * m2[7], m1[2] * m2[6] + m1[6] * m2[7], m1[3] * m2[6] + m1[7] * m2[7]}
}

func (m1 Mat4x3f) Mul3x1(m2 Vec3f) Vec4f {
	return Vec4f{m1[0] * m2[0] + m1[4] * m2[1] + m1[8] * m2[2], m1[1] * m2[0] + m1[5] * m2[1] + m1[9] * m2[2], m1[2] * m2[0] + m1[6] * m2[1] + m1[10] * m2[2], m1[3] * m2[0] + m1[7] * m2[1] + m1[11] * m2[2]}
}

func (m1 Mat4x3f) Mul3x2(m2 Mat3x2f) Mat4x2f {
	return Mat4x2f{m1[0] * m2[0] + m1[4] * m2[1] + m1[8] * m2[2], m1[1] * m2[0] + m1[5] * m2[1] + m1[9] * m2[2], m1[2] * m2[0] + m1[6] * m2[1] + m1[10] * m2[2], m1[3] * m2[0] + m1[7] * m2[1] + m1[11] * m2[2], m1[0] * m2[3] + m1[4] * m2[4] + m1[8] * m2[5], m1[1] * m2[3] + m1[5] * m2[4] + m1[9] * m2[5], m1[2] * m2[3] + m1[6] * m2[4] + m1[10] * m2[5], m1[3] * m2[3] + m1[7] * m2[4] + m1[11] * m2[5]}
}

func (m1 Mat4x3f) Mul3(m2 Mat3f) Mat4x3f {
	return Mat4x3f{m1[0] * m2[0] + m1[4] * m2[1] + m1[8] * m2[2], m1[1] * m2[0] + m1[5] * m2[1] + m1[9] * m2[2], m1[2] * m2[0] + m1[6] * m2[1] + m1[10] * m2[2], m1[3] * m2[0] + m1[7] * m2[1] + m1[11] * m2[2], m1[0] * m2[3] + m1[4] * m2[4] + m1[8] * m2[5], m1[1] * m2[3] + m1[5] * m2[4] + m1[9] * m2[5], m1[2] * m2[3] + m1[6] * m2[4] + m1[10] * m2[5], m1[3] * m2[3] + m1[7] * m2[4] + m1[11] * m2[5], m1[0] * m2[6] + m1[4] * m2[7] + m1[8] * m2[8], m1[1] * m2[6] + m1[5] * m2[7] + m1[9] * m2[8], m1[2] * m2[6] + m1[6] * m2[7] + m1[10] * m2[8], m1[3] * m2[6] + m1[7] * m2[7] + m1[11] * m2[8]}
}

func (m1 Mat4x3f) Mul3x4(m2 Mat3x4f) Mat4f {
	return Mat4f{m1[0] * m2[0] + m1[4] * m2[1] + m1[8] * m2[2], m1[1] * m2[0] + m1[5] * m2[1] + m1[9] * m2[2], m1[2] * m2[0] + m1[6] * m2[1] + m1[10] * m2[2], m1[3] * m2[0] + m1[7] * m2[1] + m1[11] * m2[2], m1[0] * m2[3] + m1[4] * m2[4] + m1[8] * m2[5], m1[1] * m2[3] + m1[5] * m2[4] + m1[9] * m2[5], m1[2] * m2[3] + m1[6] * m2[4] + m1[10] * m2[5], m1[3] * m2[3] + m1[7] * m2[4] + m1[11] * m2[5], m1[0] * m2[6] + m1[4] * m2[7] + m1[8] * m2[8], m1[1] * m2[6] + m1[5] * m2[7] + m1[9] * m2[8], m1[2] * m2[6] + m1[6] * m2[7] + m1[10] * m2[8], m1[3] * m2[6] + m1[7] * m2[7] + m1[11] * m2[8], m1[0] * m2[9] + m1[4] * m2[10] + m1[8] * m2[11], m1[1] * m2[9] + m1[5] * m2[10] + m1[9] * m2[11], m1[2] * m2[9] + m1[6] * m2[10] + m1[10] * m2[11], m1[3] * m2[9] + m1[7] * m2[10] + m1[11] * m2[11]}
}

func (m1 Mat4f) Mul4x1(m2 Vec4f) Vec4f {
	return Vec4f{m1[0] * m2[0] + m1[4] * m2[1] + m1[8] * m2[2] + m1[12] * m2[3], m1[1] * m2[0] + m1[5] * m2[1] + m1[9] * m2[2] + m1[13] * m2[3], m1[2] * m2[0] + m1[6] * m2[1] + m1[10] * m2[2] + m1[14] * m2[3], m1[3] * m2[0] + m1[7] * m2[1] + m1[11] * m2[2] + m1[15] * m2[3]}
}

func (m1 Mat4f) Mul4x2(m2 Mat4x2f) Mat4x2f {
	return Mat4x2f{m1[0] * m2[0] + m1[4] * m2[1] + m1[8] * m2[2] + m1[12] * m2[3], m1[1] * m2[0] + m1[5] * m2[1] + m1[9] * m2[2] + m1[13] * m2[3], m1[2] * m2[0] + m1[6] * m2[1] + m1[10] * m2[2] + m1[14] * m2[3], m1[3] * m2[0] + m1[7] * m2[1] + m1[11] * m2[2] + m1[15] * m2[3], m1[0] * m2[4] + m1[4] * m2[5] + m1[8] * m2[6] + m1[12] * m2[7], m1[1] * m2[4] + m1[5] * m2[5] + m1[9] * m2[6] + m1[13] * m2[7], m1[2] * m2[4] + m1[6] * m2[5] + m1[10] * m2[6] + m1[14] * m2[7], m1[3] * m2[4] + m1[7] * m2[5] + m1[11] * m2[6] + m1[15] * m2[7]}
}

func (m1 Mat4f) Mul4x3(m2 Mat4x3f) Mat4x3f {
	return Mat4x3f{m1[0] * m2[0] + m1[4] * m2[1] + m1[8] * m2[2] + m1[12] * m2[3], m1[1] * m2[0] + m1[5] * m2[1] + m1[9] * m2[2] + m1[13] * m2[3], m1[2] * m2[0] + m1[6] * m2[1] + m1[10] * m2[2] + m1[14] * m2[3], m1[3] * m2[0] + m1[7] * m2[1] + m1[11] * m2[2] + m1[15] * m2[3], m1[0] * m2[4] + m1[4] * m2[5] + m1[8] * m2[6] + m1[12] * m2[7], m1[1] * m2[4] + m1[5] * m2[5] + m1[9] * m2[6] + m1[13] * m2[7], m1[2] * m2[4] + m1[6] * m2[5] + m1[10] * m2[6] + m1[14] * m2[7], m1[3] * m2[4] + m1[7] * m2[5] + m1[11] * m2[6] + m1[15] * m2[7], m1[0] * m2[8] + m1[4] * m2[9] + m1[8] * m2[10] + m1[12] * m2[11], m1[1] * m2[8] + m1[5] * m2[9] + m1[9] * m2[10] + m1[13] * m2[11], m1[2] * m2[8] + m1[6] * m2[9] + m1[10] * m2[10] + m1[14] * m2[11], m1[3] * m2[8] + m1[7] * m2[9] + m1[11] * m2[10] + m1[15] * m2[11]}
}

func (m1 Mat4f) Mul4(m2 Mat4f) Mat4f {
	return Mat4f{m1[0] * m2[0] + m1[4] * m2[1] + m1[8] * m2[2] + m1[12] * m2[3], m1[1] * m2[0] + m1[5] * m2[1] + m1[9] * m2[2] + m1[13] * m2[3], m1[2] * m2[0] + m1[6] * m2[1] + m1[10] * m2[2] + m1[14] * m2[3], m1[3] * m2[0] + m1[7] * m2[1] + m1[11] * m2[2] + m1[15] * m2[3], m1[0] * m2[4] + m1[4] * m2[5] + m1[8] * m2[6] + m1[12] * m2[7], m1[1] * m2[4] + m1[5] * m2[5] + m1[9] * m2[6] + m1[13] * m2[7], m1[2] * m2[4] + m1[6] * m2[5] + m1[10] * m2[6] + m1[14] * m2[7], m1[3] * m2[4] + m1[7] * m2[5] + m1[11] * m2[6] + m1[15] * m2[7], m1[0] * m2[8] + m1[4] * m2[9] + m1[8] * m2[10] + m1[12] * m2[11], m1[1] * m2[8] + m1[5] * m2[9] + m1[9] * m2[10] + m1[13] * m2[11], m1[2] * m2[8] + m1[6] * m2[9] + m1[10] * m2[10] + m1[14] * m2[11], m1[3] * m2[8] + m1[7] * m2[9] + m1[11] * m2[10] + m1[15] * m2[11], m1[0] * m2[12] + m1[4] * m2[13] + m1[8] * m2[14] + m1[12] * m2[15], m1[1] * m2[12] + m1[5] * m2[13] + m1[9] * m2[14] + m1[13] * m2[15], m1[2] * m2[12] + m1[6] * m2[13] + m1[10] * m2[14] + m1[14] * m2[15], m1[3] * m2[12] + m1[7] * m2[13] + m1[11] * m2[14] + m1[15] * m2[15]}
}

func (m1 Mat2f) Transpose() Mat2f {
	return Mat2f{m1[0],m1[2],m1[1],m1[3]}
}

func (m1 Mat2x3f) Transpose() Mat3x2f {
	return Mat3x2f{m1[0],m1[3],m1[1],m1[4],m1[2],m1[5],}
}

func (m1 Mat2x4f) Transpose() Mat4x2f {
	return Mat4x2f{m1[0],m1[4],m1[1],m1[5],m1[2],m1[6],m1[3],m1[7],}
}

func (m1 Mat3x2f) Transpose() Mat2x3f {
	return Mat2x3f{m1[0],m1[2],m1[4],m1[1],m1[3],m1[5],}
}

func (m1 Mat3f) Transpose() Mat3f {
	return Mat3f{m1[0],m1[3],m1[6],m1[1],m1[4],m1[7],m1[2],m1[5],m1[8]}
}

func (m1 Mat3x4f) Transpose() Mat4x3f {
	return Mat4x3f{m1[0],m1[4],m1[8],m1[1],m1[5],m1[9],m1[2],m1[6],m1[10],m1[3],m1[7],m1[11],}
}

func (m1 Mat4x2f) Transpose() Mat2x4f {
	return Mat2x4f{m1[0],m1[2],m1[4],m1[6],m1[1],m1[3],m1[5],m1[7],}
}

func (m1 Mat4x3f) Transpose() Mat3x4f {
	return Mat3x4f{m1[0],m1[3],m1[6],m1[9],m1[1],m1[4],m1[7],m1[10],m1[2],m1[5],m1[8],m1[11],}
}

func (m1 Mat4f) Transpose() Mat4f {
	return Mat4f{m1[0],m1[4],m1[8],m1[12],m1[1],m1[5],m1[9],m1[13],m1[2],m1[6],m1[10],m1[14],m1[3],m1[7],m1[11],m1[15]}
}

func (m Mat2f) Det() float32 {
	return m[0] * m[2] - m[1] * m[3]
}

func (m Mat3f) Det() float32 {
	return m[0]*m[4]*m[8] + m[3] * m[7] * m[2] + m[6] * m[1] * m[5] - m[6] * m[4] * m[2] - m[3] * m[1] * m[8] - m[0] * m[7] * m[5]
}

func (m Mat4f) Det() float32 {
	return m[0]*m[5]*m[10]*m[15] - m[0]*m[5]*m[11]*m[14] - m[0]*m[6]*m[9]*m[15] + m[0]*m[6]*m[11]*m[13] + m[0]*m[7]*m[9]*m[14] - m[0]*m[7]*m[10]*m[13] - m[1]*m[4]*m[10]*m[15] + m[1]*m[4]*m[11]*m[14] + m[1]*m[6]*m[8]*m[15] - m[1]*m[6]*m[11]*m[12] - m[1]*m[7]*m[8]*m[14] + m[1]*m[7]*m[10]*m[12] + m[2]*m[4]*m[9]*m[15] - m[2]*m[4]*m[11]*m[13] - m[2]*m[5]*m[8]*m[15] + m[2]*m[5]*m[11]*m[12] + m[2]*m[7]*m[8]*m[13] - m[2]*m[7]*m[9]*m[12] - m[3]*m[4]*m[9]*m[14] + m[3]*m[4]*m[10]*m[13] + m[3]*m[5]*m[8]*m[14] - m[3]*m[5]*m[10]*m[12] - m[3]*m[6]*m[8]*m[13] + m[3]*m[6]*m[9]*m[12]
}

func (m Mat2f) Inv() Mat2f {
	det := m.Det()
	 if FloatEqual32(det,float32(0.0)) { 
		 return Mat2f{}
	}
	retMat := Mat2f{m[3], -m[1], -m[2], m[0]}
	 return retMat.Mul(1/det)
}

func (m Mat3f) Inv() Mat3f {
	det := m.Det()
	 if FloatEqual32(det,float32(0.0)) { 
		 return Mat3f{}
	}
	retMat := Mat3f{m[4] * m[8] -m[5] * m[7] , m[2] * m[7] -m[1] * m[8] ,m[1] * m[5] -m[2] * m[4] ,m[5] * m[6] -m[3] * m[8] ,m[0] * m[8] -m[2] * m[6] ,m[2] * m[3] -m[0] * m[5] ,m[3] * m[7] -m[4] * m[6] ,m[1] * m[6] -m[0] * m[7] ,m[0] * m[4] -m[1] * m[3]}
	 return retMat.Mul(1/det)
}

func (m Mat4f) Inv() Mat4f {
	det := m.Det()
	 if FloatEqual32(det,float32(0.0)) { 
		 return Mat4f{}
	}
	retMat := Mat4f{-m[7] * m[10] * m[13] +m[6] * m[11] * m[13] +m[7] * m[9] * m[14] -m[5] * m[11] * m[14] -m[6] * m[9] * m[15] +m[5] * m[10] * m[15] ,m[3] * m[10] * m[13] -m[2] * m[11] * m[13] -m[3] * m[9] * m[14] +m[1] * m[11] * m[14] +m[2] * m[9] * m[15] -m[1] * m[10] * m[15] ,-m[3] * m[6] * m[13] +m[2] * m[7] * m[13] +m[3] * m[5] * m[14] -m[1] * m[7] * m[14] -m[2] * m[5] * m[15] +m[1] * m[6] * m[15] ,m[3] * m[6] * m[9] -m[2] * m[7] * m[9] -m[3] * m[5] * m[10] +m[1] * m[7] * m[10] +m[2] * m[5] * m[11] -m[1] * m[6] * m[11] ,m[7] * m[10] * m[12] -m[6] * m[11] * m[12] -m[7] * m[8] * m[14] +m[4] * m[11] * m[14] +m[6] * m[8] * m[15] -m[4] * m[10] * m[15] ,-m[3] * m[10] * m[12] +m[2] * m[11] * m[12] +m[3] * m[8] * m[14] -m[0] * m[11] * m[14] -m[2] * m[8] * m[15] +m[0] * m[10] * m[15] , m[3] * m[6] * m[12] -m[2] * m[7] * m[12] -m[3] * m[4] * m[14] +m[0] * m[7] * m[14] +m[2] * m[4] * m[15] -m[0] * m[6] * m[15] ,-m[3] * m[6] * m[8] +m[2] * m[7] * m[8] +m[3] * m[4] * m[10] -m[0] * m[7] * m[10] -m[2] * m[4] * m[11] +m[0] * m[6] * m[11] ,-m[7] * m[9] * m[12] +m[5] * m[11] * m[12] +m[7] * m[8] * m[13] -m[4] * m[11] * m[13] -m[5] * m[8] * m[15] +m[4] * m[9] * m[15] ,m[3] * m[9] * m[12] -m[1] * m[11] * m[12] -m[3] * m[8] * m[13] +m[0] * m[11] * m[13] +m[1] * m[8] * m[15] -m[0] * m[9] * m[15] ,-m[3] * m[5] * m[12] +m[1] * m[7] * m[12] +m[3] * m[4] * m[13] -m[0] * m[7] * m[13] -m[1] * m[4] * m[15] +m[0] * m[5] * m[15] ,m[3] * m[5] * m[8] -m[1] * m[7] * m[8] -m[3] * m[4] * m[9] +m[0] * m[7] * m[9] +m[1] * m[4] * m[11] -m[0] * m[5] * m[11] ,m[6] * m[9] * m[12] -m[5] * m[10] * m[12] -m[6] * m[8] * m[13] +m[4] * m[10] * m[13] +m[5] * m[8] * m[14] -m[4] * m[9] * m[14] ,-m[2] * m[9] * m[12] +m[1] * m[10] * m[12] +m[2] * m[8] * m[13] -m[0] * m[10] * m[13] -m[1] * m[8] * m[14] +m[0] * m[9] * m[14] ,m[2] * m[5] * m[12] -m[1] * m[6] * m[12] -m[2] * m[4] * m[13] +m[0] * m[6] * m[13] +m[1] * m[4] * m[14] -m[0] * m[5] * m[14] ,-m[2] * m[5] * m[8] +m[1] * m[6] * m[8] +m[2] * m[4] * m[9] -m[0] * m[6] * m[9] -m[1] * m[4] * m[10] +m[0] * m[5] * m[10]}
	 return retMat.Mul(1/det)
}

func (m1 Mat2f) ApproxEqual(m2 Mat2f) bool {
	for i := range m1 {
		if !FloatEqual32(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat2x3f) ApproxEqual(m2 Mat2x3f) bool {
	for i := range m1 {
		if !FloatEqual32(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat2x4f) ApproxEqual(m2 Mat2x4f) bool {
	for i := range m1 {
		if !FloatEqual32(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3x2f) ApproxEqual(m2 Mat3x2f) bool {
	for i := range m1 {
		if !FloatEqual32(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3f) ApproxEqual(m2 Mat3f) bool {
	for i := range m1 {
		if !FloatEqual32(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3x4f) ApproxEqual(m2 Mat3x4f) bool {
	for i := range m1 {
		if !FloatEqual32(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4x2f) ApproxEqual(m2 Mat4x2f) bool {
	for i := range m1 {
		if !FloatEqual32(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4x3f) ApproxEqual(m2 Mat4x3f) bool {
	for i := range m1 {
		if !FloatEqual32(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4f) ApproxEqual(m2 Mat4f) bool {
	for i := range m1 {
		if !FloatEqual32(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat2f) ApproxEqualThreshold(m2 Mat2f, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold32(m1[i],m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat2x3f) ApproxEqualThreshold(m2 Mat2x3f, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold32(m1[i],m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat2x4f) ApproxEqualThreshold(m2 Mat2x4f, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold32(m1[i],m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat3x2f) ApproxEqualThreshold(m2 Mat3x2f, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold32(m1[i],m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat3f) ApproxEqualThreshold(m2 Mat3f, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold32(m1[i],m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat3x4f) ApproxEqualThreshold(m2 Mat3x4f, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold32(m1[i],m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat4x2f) ApproxEqualThreshold(m2 Mat4x2f, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold32(m1[i],m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat4x3f) ApproxEqualThreshold(m2 Mat4x3f, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold32(m1[i],m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat4f) ApproxEqualThreshold(m2 Mat4f, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold32(m1[i],m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat2f) ApproxFuncEqual(m2 Mat2f, eq func(float32,float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat2x3f) ApproxFuncEqual(m2 Mat2x3f, eq func(float32,float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat2x4f) ApproxFuncEqual(m2 Mat2x4f, eq func(float32,float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3x2f) ApproxFuncEqual(m2 Mat3x2f, eq func(float32,float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3f) ApproxFuncEqual(m2 Mat3f, eq func(float32,float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3x4f) ApproxFuncEqual(m2 Mat3x4f, eq func(float32,float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4x2f) ApproxFuncEqual(m2 Mat4x2f, eq func(float32,float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4x3f) ApproxFuncEqual(m2 Mat4x3f, eq func(float32,float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4f) ApproxFuncEqual(m2 Mat4f, eq func(float32,float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i],m2[i]) {
			return false
		}
	}
	return true
}

