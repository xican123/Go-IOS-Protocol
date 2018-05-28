package verifier

import (
	"testing"
	"github.com/golang/mock/gomock"
)

func TestGenesisVerify(t *testing.T) {
	Convey("Test of Genesis verify", t, func() {
		Convey("Parse Contract", func() {
			mockCtl := gomock.NewController(t)
			pool := core_mock.NewMockPool(mockCtl)
			var count int
			var k, f, k2 state.Key
			var v, v2 state.Value
			pool.EXPECT().PutHM(gomock.Any(),gomock.Any(),gomock.Any()).Times(2).Do(func(key, field state.Key, value state.Value) error {
				k ,f ,v = key, field, value
				count ++
				return nil
			})
			pool.EXPECT().Put(gomock.Any(), gomock.Any()).Do(func(key state.Key, value state.Value){
				k2, v2 = key, value
			})
			pool.EXPECT().Copy().Return(pool)
			contract := vm_mock.NewMockContract(mockCtl)
			contract.EXPECT().Code().Return(`
-- @PutHM iost abc f10000
-- @PutHM iost def f1000
-- @Put hello sworld
`)
			_, err := ParseGenesis(contract, pool)
			So(err, ShouldBeNil)
			So(count, ShouldEqual, 2)
			So(k, ShouldEqual, state.Key("iost"))
			So(v2.EncodeString(), ShouldEqual, "sworld")

		})
	})
}


func TestCacheVerifier(t *testing.T) {
	Convey("Test of CacheVerifier", t, func() {
		Convey("Verify contract", func() {
			mockCtl := gomock.NewController(t)
			pool := core_mock.NewMockPool(mockCtl)

			var k state.Key
			var v state.Value

			pool.EXPECT().Put(gomock.Any(), gomock.Any()).AnyTimes().Do(func(key state.Key, value state.Value) error {
				k = key
				v = value
				return nil
			})

			pool.EXPECT().Get(gomock.Any()).AnyTimes().Return(state.MakeVFloat(3.14), nil)

			var k2 state.Key
			var f2 state.Key
			var v2 state.Value
			pool.EXPECT().PutHM(gomock.Any(), gomock.Any(), gomock.Any()).Do(func(key, field state.Key, value state.Value) {
				k2 = key
				f2 = field
				v2 = value
			})
			v3 := state.MakeVFloat(float64(10000))
			pool.EXPECT().GetHM(gomock.Any(), gomock.Any()).Return(v3, nil)
			pool.EXPECT().Copy().AnyTimes().Return(pool)
			main := lua.NewMethod("main", 0, 1)
			code := `function main()
	a = Get("pi")
	Put("hello", a)
	return "success"
end`

		})
	})
}
