package vector_test

import (
	. "github.com/outofjungle/gotest/vector"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vector", func() {
	var (
		p1  Point
		p2  Point
		val float64
	)

	BeforeEach(func() {
		val = 0.70710678118655
	})

	Describe("Dot", func() {
		Context("When 2 unit vectors are in Q1", func() {
			It("returns positive value", func() {
				p1 = Point{X: val, Y: val}
				p2 = Point{X: val, Y: val}
				scalar := p1.Dot(p2)

				expected := 1.000000000000007
				Expect(scalar).To(Equal(expected))
			})
		})

		Context("When 1 unit vector is in Q1 and 1 in Q2", func() {
			It("returns 0", func() {
				p1 = Point{X: val, Y: val}
				p2 = Point{X: -val, Y: val}
				scalar := p1.Dot(p2)

				expected := 0.0
				Expect(scalar).To(Equal(expected))
			})
		})

		Context("When 1 unit vector is in Q1 and 1 in Q3", func() {
			It("returns positive value", func() {
				p1 = Point{X: val, Y: val}
				p2 = Point{X: -val, Y: -val}
				scalar := p1.Dot(p2)

				expected := -1.000000000000007
				Expect(scalar).To(Equal(expected))
			})
		})

		Context("When 1 unit vector is in Q1 and 1 in Q2", func() {
			It("returns 0", func() {
				p1 = Point{X: val, Y: val}
				p2 = Point{X: val, Y: -val}
				scalar := p1.Dot(p2)

				expected := 0.0
				Expect(scalar).To(Equal(expected))
			})
		})
	})
})
