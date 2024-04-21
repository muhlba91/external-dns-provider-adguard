# Changelog

## [5.0.1](https://github.com/muhlba91/external-dns-provider-adguard/compare/v5.0.0...v5.0.1) (2024-04-21)


### Miscellaneous Chores

* **ci:** update go version ([373a230](https://github.com/muhlba91/external-dns-provider-adguard/commit/373a230444c5b0b707abae63d23080278fea2509))
* **deps:** bump golang.org/x/net from 0.22.0 to 0.23.0 ([76e2f86](https://github.com/muhlba91/external-dns-provider-adguard/commit/76e2f862ba4e24ee184c0d55bcaf83784f06eca3))
* **deps:** update actions/checkout digest to 9bb5618 ([9886481](https://github.com/muhlba91/external-dns-provider-adguard/commit/9886481df80c203d1f545a0e4c6c55df4323ead1))
* **deps:** update actions/checkout digest to b4ffde6 ([121dede](https://github.com/muhlba91/external-dns-provider-adguard/commit/121dede71c3f69fc910c65d897d65174ebe5c79e))
* **deps:** update anchore/sbom-action action to v0.15.10 ([6e7d07d](https://github.com/muhlba91/external-dns-provider-adguard/commit/6e7d07d5a97c12b65f3ba831994555bd9498c2af))
* **deps:** update anchore/sbom-action action to v0.15.9 ([1179791](https://github.com/muhlba91/external-dns-provider-adguard/commit/1179791e1831fe471767ccb18bf74c262c338dd3))
* **deps:** update golang.org/x/exp digest to 93d18d7 ([789352c](https://github.com/muhlba91/external-dns-provider-adguard/commit/789352cbeb7369fc674696f2cde2758ddbe833b7))
* **deps:** update golang.org/x/exp digest to a685a6e ([358fa34](https://github.com/muhlba91/external-dns-provider-adguard/commit/358fa343e3d18b791037c488ca30efd2f2f48b21))
* **deps:** update golang.org/x/exp digest to a85f2c6 ([6cd3cf8](https://github.com/muhlba91/external-dns-provider-adguard/commit/6cd3cf8d489c15e10c929a138589608de361b966))
* **deps:** update golang.org/x/exp digest to c0f41cb ([1e33d3f](https://github.com/muhlba91/external-dns-provider-adguard/commit/1e33d3f9fc4e5498db99922b3a67d0934688948e))
* **deps:** update golang.org/x/exp digest to c7f7c64 ([daf0566](https://github.com/muhlba91/external-dns-provider-adguard/commit/daf056689774f683533ca9e6b724510977f77049))
* **deps:** update golang.org/x/exp digest to fe59bbe ([2b1d7d5](https://github.com/muhlba91/external-dns-provider-adguard/commit/2b1d7d58063719b9882e5b3002c320f526f939db))
* **deps:** update module github.com/caarlos0/env/v10 to v11 ([1dee616](https://github.com/muhlba91/external-dns-provider-adguard/commit/1dee6165545995c7c379f77b1f3c2a3bba91c923))
* **deps:** update module github.com/caarlos0/env/v10 to v11 ([c624a11](https://github.com/muhlba91/external-dns-provider-adguard/commit/c624a116e88a550df884a009a2eaa2773e7e304f))
* **deps:** update module sigs.k8s.io/external-dns to v0.14.1 ([4423371](https://github.com/muhlba91/external-dns-provider-adguard/commit/4423371450e233dbdb668818237933aa28518750))
* **deps:** update sigstore/cosign-installer action to v3.5.0 ([e97c3bc](https://github.com/muhlba91/external-dns-provider-adguard/commit/e97c3bc44777a634da3230de6e162120a2495928))

## [5.0.0](https://github.com/muhlba91/external-dns-provider-adguard/compare/v4.0.0...v5.0.0) (2024-03-02)


### ⚠ BREAKING CHANGES

* removes the migration path to the new rules syntax

### Features

* deprecate migration from the old rules syntax to the new syntax ([4afa607](https://github.com/muhlba91/external-dns-provider-adguard/commit/4afa6071ede1327eafaeee975bcf855a53b0cb9e))

## [4.0.0](https://github.com/muhlba91/external-dns-provider-adguard/compare/v3.1.0...v4.0.0) (2024-03-02)


### ⚠ BREAKING CHANGES

* new rules syntax - existing rules will be converted

### Bug Fixes

* fixes subdomain handling by introducing a new rules syntax ([d3bf273](https://github.com/muhlba91/external-dns-provider-adguard/commit/d3bf2734223e4c4c939f651c81d4b4a52f8fc12e))


### Miscellaneous Chores

* **ci:** adopt release please for v4 ([05b5262](https://github.com/muhlba91/external-dns-provider-adguard/commit/05b5262f0039bc51b27dac57909eb65c5688c175))
* **deps:** update actions/setup-go action to v5 ([53258a3](https://github.com/muhlba91/external-dns-provider-adguard/commit/53258a3f0813d73fb13d7a6f47fd63303835e12d))
* **deps:** update anchore/sbom-action action to v0.15.0 ([4f3cd3e](https://github.com/muhlba91/external-dns-provider-adguard/commit/4f3cd3ea1cd27f4cb235ef0d2852464ea03b07fd))
* **deps:** update anchore/sbom-action action to v0.15.1 ([923a489](https://github.com/muhlba91/external-dns-provider-adguard/commit/923a489321c704b736d291131c495e777c9074ab))
* **deps:** update anchore/sbom-action action to v0.15.2 ([ab289c9](https://github.com/muhlba91/external-dns-provider-adguard/commit/ab289c96f49fc4e9347fbc0abd5e783f3248a641))
* **deps:** update anchore/sbom-action action to v0.15.3 ([16e5b02](https://github.com/muhlba91/external-dns-provider-adguard/commit/16e5b02e74cc0583c7ad2eaa332d4b29d1f6612c))
* **deps:** update anchore/sbom-action action to v0.15.4 ([2545229](https://github.com/muhlba91/external-dns-provider-adguard/commit/25452290217299c19dce399ec7035fab29f8d941))
* **deps:** update anchore/sbom-action action to v0.15.5 ([9915791](https://github.com/muhlba91/external-dns-provider-adguard/commit/9915791d8c1435ada4d5491ebd8f3b0d2322bd5c))
* **deps:** update anchore/sbom-action action to v0.15.6 ([141b8b0](https://github.com/muhlba91/external-dns-provider-adguard/commit/141b8b063b54184f696ad19282def326920209ac))
* **deps:** update anchore/sbom-action action to v0.15.7 ([3cd37b5](https://github.com/muhlba91/external-dns-provider-adguard/commit/3cd37b516bbc032d892d9807eaa4e1540cb2f18c))
* **deps:** update anchore/sbom-action action to v0.15.8 ([0783b32](https://github.com/muhlba91/external-dns-provider-adguard/commit/0783b32dc0314055830d458bb3026ecf346d01bb))
* **deps:** update golang.org/x/exp digest to 02704c9 ([0525b07](https://github.com/muhlba91/external-dns-provider-adguard/commit/0525b0706dd026356a2abc122d52ee42669c2537))
* **deps:** update golang.org/x/exp digest to 0dcbfd6 ([5bc7a38](https://github.com/muhlba91/external-dns-provider-adguard/commit/5bc7a38fc9a69aa18516103acc6d3297d7a50dfa))
* **deps:** update golang.org/x/exp digest to 1b97071 ([f52c549](https://github.com/muhlba91/external-dns-provider-adguard/commit/f52c54971e091a459df1acaa02c08e930aeb0d96))
* **deps:** update golang.org/x/exp digest to 2c58cdc ([4754108](https://github.com/muhlba91/external-dns-provider-adguard/commit/47541080949f3342920e542ef6bd24b384c128bc))
* **deps:** update golang.org/x/exp digest to 6522937 ([de71206](https://github.com/muhlba91/external-dns-provider-adguard/commit/de712068b9b82d9b7b8439c11011541ebb1c4b89))
* **deps:** update golang.org/x/exp digest to 73b9e39 ([49111d1](https://github.com/muhlba91/external-dns-provider-adguard/commit/49111d1badc21ad9f00eaf25cef04e4d8064b9ae))
* **deps:** update golang.org/x/exp digest to 814bf88 ([5eb7e15](https://github.com/muhlba91/external-dns-provider-adguard/commit/5eb7e1532d2d0573b2c8d9bbd36f41094feb3fe4))
* **deps:** update golang.org/x/exp digest to aacd6d4 ([d6968a0](https://github.com/muhlba91/external-dns-provider-adguard/commit/d6968a0eb79ad700fd2be8f4bc88ac6d75424a7e))
* **deps:** update golang.org/x/exp digest to be819d1 ([8b4bef0](https://github.com/muhlba91/external-dns-provider-adguard/commit/8b4bef0e6d3c38950f42e3e6579fc9c83868bb6a))
* **deps:** update golang.org/x/exp digest to db7319d ([08077cc](https://github.com/muhlba91/external-dns-provider-adguard/commit/08077cc4fcfe1aac30a4276345cb211bf8be1d51))
* **deps:** update golang.org/x/exp digest to dc181d7 ([2e192c1](https://github.com/muhlba91/external-dns-provider-adguard/commit/2e192c14dc9079572e1f9b87d7adcffab61a7c3e))
* **deps:** update golang.org/x/exp digest to ec58324 ([d197313](https://github.com/muhlba91/external-dns-provider-adguard/commit/d19731334e7a64f1930f2d4b67d19747f1bffd32))
* **deps:** update golang.org/x/exp digest to f3f8817 ([433160f](https://github.com/muhlba91/external-dns-provider-adguard/commit/433160f63c5ccee5d56bd319901b2c9df86d6969))
* **deps:** update golangci/golangci-lint-action action to v4 ([a150d20](https://github.com/muhlba91/external-dns-provider-adguard/commit/a150d20d2735749c530fd770d691f22d2e534275))
* **deps:** update google-github-actions/release-please-action action to v4 ([aa61995](https://github.com/muhlba91/external-dns-provider-adguard/commit/aa61995b64ca00d01993785a061bdfd9f0321aa3))
* **deps:** update module github.com/go-chi/chi/v5 to v5.0.11 ([51209c2](https://github.com/muhlba91/external-dns-provider-adguard/commit/51209c266b5b2abd3d776468d05532ff38e5460f))
* **deps:** update module github.com/go-chi/chi/v5 to v5.0.12 ([a47726d](https://github.com/muhlba91/external-dns-provider-adguard/commit/a47726dc3e13245ad0f2e63fdb5ad90b0c1e5450))
* **deps:** update module github.com/stretchr/testify to v1.9.0 ([fed5a11](https://github.com/muhlba91/external-dns-provider-adguard/commit/fed5a1155d7e8a987710fa84d0470aa102ef941d))
* **deps:** update sigstore/cosign-installer action to v3.3.0 ([520316d](https://github.com/muhlba91/external-dns-provider-adguard/commit/520316d9626c54c3b0b2a33c001b0c7adee6e1e6))
* **deps:** update sigstore/cosign-installer action to v3.4.0 ([df22d7e](https://github.com/muhlba91/external-dns-provider-adguard/commit/df22d7e72bf609dcda43571d4bd30ffe03d09315))

## [3.1.0](https://github.com/muhlba91/external-dns-provider-adguard/compare/v3.0.0...v3.1.0) (2023-11-11)


### Features

* change health endpoint to /healthz ([29e9033](https://github.com/muhlba91/external-dns-provider-adguard/commit/29e903324a7e698797cb50b2e60ed681577503d6))


### Miscellaneous Chores

* **deps:** update golang.org/x/exp digest to 9a3e603 ([d2aaee5](https://github.com/muhlba91/external-dns-provider-adguard/commit/d2aaee5c57827876ff635320b9e2936dfdf5b383))

## [3.0.0](https://github.com/muhlba91/external-dns-provider-adguard/compare/v2.0.0...v3.0.0) (2023-11-09)


### Bug Fixes

* fix rule creation ([2c489d4](https://github.com/muhlba91/external-dns-provider-adguard/commit/2c489d45860ec0737d69d716882cdbe6355f9ee7))

## [2.0.0](https://github.com/muhlba91/external-dns-provider-adguard/compare/v1.0.1...v2.0.0) (2023-11-09)


### Bug Fixes

* fix release docker images ([593192a](https://github.com/muhlba91/external-dns-provider-adguard/commit/593192aec9d3448ec5fbce0a1039936550b609ad))

## [1.0.1](https://github.com/muhlba91/external-dns-provider-adguard/compare/v1.0.0...v1.0.1) (2023-11-09)


### Bug Fixes

* fix release permissions ([2c16420](https://github.com/muhlba91/external-dns-provider-adguard/commit/2c16420114023fca85c6f5fdd31080799db6fbe5))

## [1.0.0](https://github.com/muhlba91/external-dns-provider-adguard/compare/v0.1.0...v1.0.0) (2023-11-09)


### Features

* add arm v7 support ([44e6e5d](https://github.com/muhlba91/external-dns-provider-adguard/commit/44e6e5d9d134e22959d62af1cf7bb0b2243421ad))
* initial version ([314a602](https://github.com/muhlba91/external-dns-provider-adguard/commit/314a602080d82bba6ede2ff1b0fa0c165303470c))
* support all external-dns records; fix [#12](https://github.com/muhlba91/external-dns-provider-adguard/issues/12) ([b455a9d](https://github.com/muhlba91/external-dns-provider-adguard/commit/b455a9dc6e3d1cccca2c8f2beb5c3c7c71021fe4))
* support arbitrary TXT records; fix [#13](https://github.com/muhlba91/external-dns-provider-adguard/issues/13) ([f2090b6](https://github.com/muhlba91/external-dns-provider-adguard/commit/f2090b6d1ed40bf167298f8f52d6e9afc05528e0))


### Bug Fixes

* initial commit ([fd0f1e3](https://github.com/muhlba91/external-dns-provider-adguard/commit/fd0f1e329a25ebf5996524b519a4c0ce84ecda1d))


### Miscellaneous Chores

* **deps:** update actions/checkout digest to b4ffde6 ([69b0901](https://github.com/muhlba91/external-dns-provider-adguard/commit/69b090183a091844785edb920c53ecfff3df487f))
* **deps:** update env to v10 ([a6523eb](https://github.com/muhlba91/external-dns-provider-adguard/commit/a6523eb63153feb9681c105535d79e176044b5b6))
* **deps:** update golang.org/x/exp digest to 2478ac8 ([d955d8d](https://github.com/muhlba91/external-dns-provider-adguard/commit/d955d8dd218f2a1dcd3af0708ee078371ebe6c7a))
* **deps:** update module github.com/caarlos0/env/v8 to v10 ([0785890](https://github.com/muhlba91/external-dns-provider-adguard/commit/078589021f3ffcc466c7f9964a082110203d457e))
* **deps:** update module github.com/go-chi/chi/v5 to v5.0.10 ([027cc50](https://github.com/muhlba91/external-dns-provider-adguard/commit/027cc50972bf08d867c4f8699200ca0a40c03786))
* **deps:** update module sigs.k8s.io/external-dns to v0.14.0 ([fd0a900](https://github.com/muhlba91/external-dns-provider-adguard/commit/fd0a900843f32cd43048f7cea036483fa13284cf))
* **deps:** update sigs.k8s.io/external-dns digest to 9f7167f ([afdd366](https://github.com/muhlba91/external-dns-provider-adguard/commit/afdd366b74e5f417006406e43b78f034a181d139))
* **deps:** update sigstore/cosign-installer action to v3.2.0 ([7631196](https://github.com/muhlba91/external-dns-provider-adguard/commit/76311961c0d786e2a6a2fcc9c489dbe5dba68d42))
* release 1.0.0 ([fc6232d](https://github.com/muhlba91/external-dns-provider-adguard/commit/fc6232d3f51f8dacc073cc93327cdb9d9c606508))
