// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "205bea690c8a769e008ee539a20e1127"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"bda96d7cf0c7c7c8b75f7a549d4c9dba": "1f8b08000000000000ffb4576f6fdbb6137edf4f7175fb031220946437891355317e43d762038ab618ba177b498b6789094d0a2415db33fcdd0752ff6d2568b03548628a3c3e777ceeb9a39cbcfef5eb87ef7f7dfb08b95d8bc5aba4fa004872a4cc0d0012cbadc0c5b75fbefcfee135500bfb7df0e71f9f0f07206ef851ebc321092ba36a83e0f2c18f00348abb89b13b812647b4937a3ad7b8ba9be4d616260ec394c97b13a442956c25a8c62055eb90ded36d28f8d28439cf72c1b3dc06f7269c46c13498861564f868823597416a4c851cfe9c082ce562c3254b8d09a7c16570ddce8cbb37a9e68505a3d37f75c86ece79b93793451256d0ffad1f416556d20c4d98a99fe1eade10c1251259ae97a88df33c0b6e82e8698b678358342914f726e092dbdf1a182eb3aff2b3a2ececfcfdb1d1672ef14b053fb439f6e1a4027657e0ddc4e2d6862ebb8dcbff3fe06ea5e91a0db4b1c3be5e0488fed77b0058d2f421d3aa942c069d2de9d96c165dc0bbfa2f0adeb551021cdad1347a21cc7404e6f0aa1e141a2fea61aa18f680574a5a62f8df184314dcccaf70ddc16c726e919882a6183b08b2d1b4e896d523ea95501bb28d412a89ef8fbd068ef69eaf8232c6651643346e4a84ec596f38b379ec8978d29cc85178b829b64fbbb03a96362769ce053bbb3eefc01c35173f6c5dabb417c04b5244255f53cb958c6166c0c99e6a981ae072e5748a4085452da9c54e63a747b29aa6f8d293fde8a6d303f6827e3ae1e35bc906970fdc12abca342729154295760803501ad4c4a0c0f464c91521a18267328614a545ddada54a281dc39bdbdbdb6e72a934434db4232e8669b105a30467f0663e9f77568fa82d4fa968a0adea09bc1654837135222acb061c9faa91085c39ffd1338af45be3dcd5d3736a721aaa7e83e8eabc4383aa7db98ee52fedb0b9b593a5623b480535e66eb2cc48a6e98e4ca3a86e64fe7647ddf485d64c2323d7515451ee7b00145b7203c58e5c36b42ebb21a32647d6f28d8ccca3a8ba099b8699e4d306df83ceb6a26da6bed96a25b3eaddc21dc43ff5960b2a07db9d8fdb264007d57b1d494267de3a0ef36917c46c8022b2c9a27b73c967352961c5caa24a52c2f8e31135f32135be812e95602d4933586aa40f64a334eb2e8dc131d69a5c3aef675c32dc42f0a9ba4da2f3e05329d39353f88731eb6fd4e687433cb6e46eba21501232fed83bd730f12b815b70ff48aa046c45ecc75a6d9accda97a5bc47dc86ac4ae13137641ace609955dcf535d0333739656a43b894a87b260049a1b1ffece77cd9d53b333501462d7535652cd596acb45adf4df67b532e619c23b83a1c5c2aea4537652072cc39e4a1bb241c4450133a78e81dbfeb4fb5ea7cdf3d26c3957801aaa029b73b721581d5541aeeba2ca9678195dab75df22e8a00a941c22551a505df33e266f3348a60405e1dc138cd9d70eb7280b22850a7d420142ed9c5965c4f16df5dd4f1e0a82765d1341628c835ace996e4c4a41a5176ef083b424bab86e92cc590defd5e539921bcf5b9b880b7fe250be2bb26678703ecf77c05b2b171891a6648f0b6c65c8faad89e9ca8a69f9ea33415e4b261d1ac6b8d3752afae90fac19fda65c4376964debc91f611fa7100430a7d02dcee91bef1d49e367d4ebc15536df318e4ea1997fe04f3f64238fe6971ab3673e15f53bad9aac38cf9190fe0a9e991aaae5786b55df77bf35c8df7421b16f6dbbab22bdd8c977713e6589b398d3d09053fd62f4ae635ea3ffb520ffb5a1feb1b836112ba9b7bf12a09ab6fe2ff040000ffffe4a8b4faa10f0000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("recovr", "./templates")
		b.SetResolver("panic.html", packr.Pointer{ForwardBox: gk, ForwardPath: "bda96d7cf0c7c7c8b75f7a549d4c9dba"})
	}()
	return nil
}()