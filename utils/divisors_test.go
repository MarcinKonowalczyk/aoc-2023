package utils

import (
	"testing"
)

type testCase struct {
	n        int
	divisors []int
}

func TestDivisorsUpTo100(t *testing.T) {
	test_cases := []testCase{
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 3}},
		{4, []int{1, 2, 4}},
		{5, []int{1, 5}},
		{6, []int{1, 2, 3, 6}},
		{7, []int{1, 7}},
		{8, []int{1, 2, 4, 8}},
		{9, []int{1, 3, 9}},
		{10, []int{1, 2, 5, 10}},
		{11, []int{1, 11}},
		{12, []int{1, 2, 3, 4, 6, 12}},
		{13, []int{1, 13}},
		{14, []int{1, 2, 7, 14}},
		{15, []int{1, 3, 5, 15}},
		{16, []int{1, 2, 4, 8, 16}},
		{17, []int{1, 17}},
		{18, []int{1, 2, 3, 6, 9, 18}},
		{19, []int{1, 19}},
		{20, []int{1, 2, 4, 5, 10, 20}},
		{21, []int{1, 3, 7, 21}},
		{22, []int{1, 2, 11, 22}},
		{23, []int{1, 23}},
		{24, []int{1, 2, 3, 4, 6, 8, 12, 24}},
		{25, []int{1, 5, 25}},
		{26, []int{1, 2, 13, 26}},
		{27, []int{1, 3, 9, 27}},
		{28, []int{1, 2, 4, 7, 14, 28}},
		{29, []int{1, 29}},
		{30, []int{1, 2, 3, 5, 6, 10, 15, 30}},
		{31, []int{1, 31}},
		{32, []int{1, 2, 4, 8, 16, 32}},
		{33, []int{1, 3, 11, 33}},
		{34, []int{1, 2, 17, 34}},
		{35, []int{1, 5, 7, 35}},
		{36, []int{1, 2, 3, 4, 6, 9, 12, 18, 36}},
		{37, []int{1, 37}},
		{38, []int{1, 2, 19, 38}},
		{39, []int{1, 3, 13, 39}},
		{40, []int{1, 2, 4, 5, 8, 10, 20, 40}},
		{41, []int{1, 41}},
		{42, []int{1, 2, 3, 6, 7, 14, 21, 42}},
		{43, []int{1, 43}},
		{44, []int{1, 2, 4, 11, 22, 44}},
		{45, []int{1, 3, 5, 9, 15, 45}},
		{46, []int{1, 2, 23, 46}},
		{47, []int{1, 47}},
		{48, []int{1, 2, 3, 4, 6, 8, 12, 16, 24, 48}},
		{49, []int{1, 7, 49}},
		{50, []int{1, 2, 5, 10, 25, 50}},
		{51, []int{1, 3, 17, 51}},
		{52, []int{1, 2, 4, 13, 26, 52}},
		{53, []int{1, 53}},
		{54, []int{1, 2, 3, 6, 9, 18, 27, 54}},
		{55, []int{1, 5, 11, 55}},
		{56, []int{1, 2, 4, 7, 8, 14, 28, 56}},
		{57, []int{1, 3, 19, 57}},
		{58, []int{1, 2, 29, 58}},
		{59, []int{1, 59}},
		{60, []int{1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30, 60}},
		{61, []int{1, 61}},
		{62, []int{1, 2, 31, 62}},
		{63, []int{1, 3, 7, 9, 21, 63}},
		{64, []int{1, 2, 4, 8, 16, 32, 64}},
		{65, []int{1, 5, 13, 65}},
		{66, []int{1, 2, 3, 6, 11, 22, 33, 66}},
		{67, []int{1, 67}},
		{68, []int{1, 2, 4, 17, 34, 68}},
		{69, []int{1, 3, 23, 69}},
		{70, []int{1, 2, 5, 7, 10, 14, 35, 70}},
		{71, []int{1, 71}},
		{72, []int{1, 2, 3, 4, 6, 8, 9, 12, 18, 24, 36, 72}},
		{73, []int{1, 73}},
		{74, []int{1, 2, 37, 74}},
		{75, []int{1, 3, 5, 15, 25, 75}},
		{76, []int{1, 2, 4, 19, 38, 76}},
		{77, []int{1, 7, 11, 77}},
		{78, []int{1, 2, 3, 6, 13, 26, 39, 78}},
		{79, []int{1, 79}},
		{80, []int{1, 2, 4, 5, 8, 10, 16, 20, 40, 80}},
		{81, []int{1, 3, 9, 27, 81}},
		{82, []int{1, 2, 41, 82}},
		{83, []int{1, 83}},
		{84, []int{1, 2, 3, 4, 6, 7, 12, 14, 21, 28, 42, 84}},
		{85, []int{1, 5, 17, 85}},
		{86, []int{1, 2, 43, 86}},
		{87, []int{1, 3, 29, 87}},
		{88, []int{1, 2, 4, 8, 11, 22, 44, 88}},
		{89, []int{1, 89}},
		{90, []int{1, 2, 3, 5, 6, 9, 10, 15, 18, 30, 45, 90}},
		{91, []int{1, 7, 13, 91}},
		{92, []int{1, 2, 4, 23, 46, 92}},
		{93, []int{1, 3, 31, 93}},
		{94, []int{1, 2, 47, 94}},
		{95, []int{1, 5, 19, 95}},
		{96, []int{1, 2, 3, 4, 6, 8, 12, 16, 24, 32, 48, 96}},
		{97, []int{1, 97}},
		{98, []int{1, 2, 7, 14, 49, 98}},
		{99, []int{1, 3, 9, 11, 33, 99}},
		{100, []int{1, 2, 4, 5, 10, 20, 25, 50, 100}},
		{101, []int{1, 101}},
	}
	for _, tc := range test_cases {
		divisors := Divisors(tc.n)
		AssertEqualArrays(t, divisors, tc.divisors)
	}
}

func TestDivisorsZero(t *testing.T) {
	divisors := Divisors(0)
	expected := []int{}
	AssertEqualArrays(t, divisors, expected)
}

func TestDivisorsNegative(t *testing.T) {
	divisors := Divisors(-1)
	expected := []int{}
	AssertEqualArrays(t, divisors, expected)
}

func TestDivisorsLarge(t *testing.T) {
	test_cases := []testCase{
		{11088, []int{1, 2, 3, 4, 6, 7, 8, 9, 11, 12, 14, 16, 18, 21, 22, 24, 28, 33, 36, 42, 44, 48, 56, 63, 66, 72, 77, 84, 88, 99, 112, 126, 132, 144, 154, 168, 176, 198, 231, 252, 264, 308, 336, 396, 462, 504, 528, 616, 693, 792, 924, 1008, 1232, 1386, 1584, 1848, 2772, 3696, 5544, 11088}},
		{1153152, []int{1, 2, 3, 4, 6, 7, 8, 9, 11, 12, 13, 14, 16, 18, 21, 22, 24, 26, 28, 32, 33, 36, 39, 42, 44, 48, 52, 56, 63, 64, 66, 72, 77, 78, 84, 88, 91, 96, 99, 104, 112, 117, 126, 128, 132, 143, 144, 154, 156, 168, 176, 182, 192, 198, 208, 224, 231, 234, 252, 264, 273, 286, 288, 308, 312, 336, 352, 364, 384, 396, 416, 429, 448, 462, 468, 504, 528, 546, 572, 576, 616, 624, 672, 693, 704, 728, 792, 819, 832, 858, 896, 924, 936, 1001, 1008, 1056, 1092, 1144, 1152, 1232, 1248, 1287, 1344, 1386, 1408, 1456, 1584, 1638, 1664, 1716, 1848, 1872, 2002, 2016, 2112, 2184, 2288, 2464, 2496, 2574, 2688, 2772, 2912, 3003, 3168, 3276, 3432, 3696, 3744, 4004, 4032, 4224, 4368, 4576, 4928, 4992, 5148, 5544, 5824, 6006, 6336, 6552, 6864, 7392, 7488, 8008, 8064, 8736, 9009, 9152, 9856, 10296, 11088, 11648, 12012, 12672, 13104, 13728, 14784, 14976, 16016, 17472, 18018, 18304, 20592, 22176, 24024, 26208, 27456, 29568, 32032, 34944, 36036, 41184, 44352, 48048, 52416, 54912, 64064, 72072, 82368, 88704, 96096, 104832, 128128, 144144, 164736, 192192, 288288, 384384, 576576, 1153152}},
		{72648576, []int{1, 2, 3, 4, 6, 7, 8, 9, 11, 12, 13, 14, 16, 18, 21, 22, 24, 26, 27, 28, 32, 33, 36, 39, 42, 44, 48, 49, 52, 54, 56, 63, 64, 66, 72, 77, 78, 81, 84, 88, 91, 96, 98, 99, 104, 108, 112, 117, 126, 128, 132, 143, 144, 147, 154, 156, 162, 168, 176, 182, 189, 192, 196, 198, 208, 216, 224, 231, 234, 252, 264, 273, 286, 288, 294, 297, 308, 312, 324, 336, 351, 352, 364, 378, 384, 392, 396, 416, 429, 432, 441, 448, 462, 468, 504, 528, 539, 546, 567, 572, 576, 588, 594, 616, 624, 637, 648, 672, 693, 702, 704, 728, 756, 784, 792, 819, 832, 858, 864, 882, 891, 896, 924, 936, 1001, 1008, 1053, 1056, 1078, 1092, 1134, 1144, 1152, 1176, 1188, 1232, 1248, 1274, 1287, 1296, 1323, 1344, 1386, 1404, 1408, 1456, 1512, 1568, 1584, 1617, 1638, 1664, 1716, 1728, 1764, 1782, 1848, 1872, 1911, 2002, 2016, 2079, 2106, 2112, 2156, 2184, 2268, 2288, 2352, 2376, 2457, 2464, 2496, 2548, 2574, 2592, 2646, 2688, 2772, 2808, 2912, 3003, 3024, 3136, 3168, 3234, 3276, 3432, 3456, 3528, 3564, 3696, 3744, 3822, 3861, 3969, 4004, 4032, 4158, 4212, 4224, 4312, 4368, 4536, 4576, 4704, 4752, 4851, 4914, 4928, 4992, 5096, 5148, 5184, 5292, 5544, 5616, 5733, 5824, 6006, 6048, 6237, 6272, 6336, 6468, 6552, 6864, 7007, 7056, 7128, 7371, 7392, 7488, 7644, 7722, 7938, 8008, 8064, 8316, 8424, 8624, 8736, 9009, 9072, 9152, 9408, 9504, 9702, 9828, 9856, 10192, 10296, 10368, 10584, 11088, 11232, 11466, 11583, 11648, 12012, 12096, 12474, 12672, 12936, 13104, 13728, 14014, 14112, 14256, 14553, 14742, 14784, 14976, 15288, 15444, 15876, 16016, 16632, 16848, 17199, 17248, 17472, 18018, 18144, 18304, 18816, 19008, 19404, 19656, 20384, 20592, 21021, 21168, 22176, 22464, 22932, 23166, 24024, 24192, 24948, 25872, 26208, 27027, 27456, 28028, 28224, 28512, 29106, 29484, 29568, 30576, 30888, 31752, 32032, 33264, 33696, 34398, 34496, 34944, 36036, 36288, 38016, 38808, 39312, 40768, 41184, 42042, 42336, 43659, 44352, 44928, 45864, 46332, 48048, 49896, 51597, 51744, 52416, 54054, 54912, 56056, 56448, 57024, 58212, 58968, 61152, 61776, 63063, 63504, 64064, 66528, 67392, 68796, 68992, 72072, 72576, 77616, 78624, 81081, 81536, 82368, 84084, 84672, 87318, 88704, 91728, 92664, 96096, 99792, 103194, 103488, 104832, 108108, 112112, 114048, 116424, 117936, 122304, 123552, 126126, 127008, 128128, 133056, 134784, 137592, 144144, 155232, 157248, 162162, 164736, 168168, 169344, 174636, 183456, 185328, 189189, 192192, 199584, 206388, 206976, 216216, 224224, 232848, 235872, 244608, 247104, 252252, 254016, 266112, 275184, 288288, 310464, 314496, 324324, 336336, 349272, 366912, 370656, 378378, 384384, 399168, 412776, 432432, 448448, 465696, 471744, 494208, 504504, 508032, 550368, 567567, 576576, 620928, 648648, 672672, 698544, 733824, 741312, 756756, 798336, 825552, 864864, 896896, 931392, 943488, 1009008, 1100736, 1135134, 1153152, 1297296, 1345344, 1397088, 1482624, 1513512, 1651104, 1729728, 1862784, 2018016, 2201472, 2270268, 2594592, 2690688, 2794176, 3027024, 3302208, 3459456, 4036032, 4540536, 5189184, 5588352, 6054048, 6604416, 8072064, 9081072, 10378368, 12108096, 18162144, 24216192, 36324288, 72648576}},
		{7513606324224, []int{1, 2, 3, 4, 6, 7, 8, 9, 11, 12, 13, 14, 16, 18, 21, 22, 24, 26, 27, 28, 32, 33, 36, 39, 42, 44, 48, 49, 52, 54, 56, 63, 64, 66, 72, 77, 78, 81, 84, 88, 91, 96, 98, 99, 101, 104, 108, 112, 117, 126, 128, 132, 143, 144, 147, 154, 156, 162, 168, 176, 182, 189, 192, 196, 198, 202, 208, 216, 224, 231, 234, 252, 256, 264, 273, 286, 288, 294, 297, 303, 308, 312, 324, 336, 351, 352, 364, 378, 384, 392, 396, 404, 416, 429, 432, 441, 448, 462, 468, 504, 512, 528, 539, 546, 567, 572, 576, 588, 594, 606, 616, 624, 637, 648, 672, 693, 702, 704, 707, 728, 756, 768, 784, 792, 808, 819, 832, 858, 864, 882, 891, 896, 909, 924, 936, 1001, 1008, 1024, 1053, 1056, 1078, 1092, 1111, 1134, 1144, 1152, 1176, 1188, 1212, 1232, 1248, 1274, 1287, 1296, 1313, 1323, 1344, 1386, 1404, 1408, 1414, 1456, 1512, 1536, 1568, 1584, 1616, 1617, 1638, 1664, 1716, 1728, 1764, 1782, 1792, 1818, 1848, 1872, 1911, 2002, 2016, 2048, 2079, 2106, 2112, 2121, 2156, 2184, 2222, 2268, 2288, 2304, 2352, 2376, 2424, 2457, 2464, 2496, 2548, 2574, 2592, 2626, 2646, 2688, 2727, 2772, 2808, 2816, 2828, 2912, 3003, 3024, 3072, 3136, 3168, 3232, 3234, 3276, 3328, 3333, 3432, 3456, 3528, 3564, 3584, 3636, 3696, 3744, 3822, 3861, 3939, 3969, 4004, 4032, 4096, 4158, 4212, 4224, 4242, 4312, 4368, 4444, 4536, 4576, 4608, 4704, 4752, 4848, 4851, 4914, 4928, 4949, 4992, 5096, 5148, 5184, 5252, 5292, 5376, 5454, 5544, 5616, 5632, 5656, 5733, 5824, 6006, 6048, 6144, 6237, 6272, 6336, 6363, 6464, 6468, 6552, 6656, 6666, 6864, 6912, 7007, 7056, 7128, 7168, 7272, 7371, 7392, 7488, 7644, 7722, 7777, 7878, 7938, 8008, 8064, 8181, 8192, 8316, 8424, 8448, 8484, 8624, 8736, 8888, 9009, 9072, 9152, 9191, 9216, 9408, 9504, 9696, 9702, 9828, 9856, 9898, 9984, 9999, 10192, 10296, 10368, 10504, 10584, 10752, 10908, 11088, 11232, 11264, 11312, 11466, 11583, 11648, 11817, 12012, 12096, 12288, 12474, 12544, 12672, 12726, 12928, 12936, 13104, 13312, 13332, 13728, 13824, 14014, 14112, 14256, 14336, 14443, 14544, 14553, 14742, 14784, 14847, 14976, 15288, 15444, 15554, 15756, 15876, 16016, 16128, 16362, 16384, 16632, 16848, 16896, 16968, 17199, 17248, 17472, 17776, 18018, 18144, 18304, 18382, 18432, 18816, 19008, 19089, 19392, 19404, 19656, 19712, 19796, 19968, 19998, 20384, 20592, 20736, 21008, 21021, 21168, 21504, 21816, 22176, 22464, 22528, 22624, 22932, 23166, 23296, 23331, 23634, 24024, 24192, 24576, 24948, 25088, 25344, 25452, 25856, 25872, 26208, 26624, 26664, 27027, 27456, 27573, 27648, 28028, 28224, 28512, 28672, 28886, 29088, 29106, 29484, 29568, 29694, 29952, 29997, 30576, 30888, 31108, 31512, 31752, 32032, 32256, 32724, 32768, 33264, 33696, 33792, 33936, 34398, 34496, 34944, 35451, 35552, 36036, 36288, 36608, 36764, 36864, 37632, 38016, 38178, 38784, 38808, 39312, 39424, 39592, 39936, 39996, 40768, 41184, 41472, 42016, 42042, 42336, 43008, 43329, 43632, 43659, 44352, 44541, 44928, 45056, 45248, 45864, 46332, 46592, 46662, 47268, 48048, 48384, 49152, 49896, 50176, 50688, 50904, 51597, 51712, 51744, 52416, 53248, 53328, 54054, 54439, 54912, 55146, 55296, 56056, 56448, 57024, 57267, 57344, 57772, 58176, 58212, 58968, 59136, 59388, 59904, 59994, 61152, 61776, 62216, 63024, 63063, 63504, 64064, 64337, 64512, 65448, 65536, 66528, 67392, 67584, 67872, 68796, 68992, 69888, 69993, 70902, 71104, 72072, 72576, 73216, 73528, 73728, 75264, 76032, 76356, 77568, 77616, 78624, 78848, 79184, 79872, 79992, 81081, 81536, 82368, 82719, 82944, 84032, 84084, 84672, 86016, 86658, 87264, 87318, 88704, 89082, 89856, 89991, 90112, 90496, 91728, 92664, 93184, 93324, 94536, 96096, 96768, 98304, 99792, 100352, 101101, 101376, 101808, 103194, 103424, 103488, 104832, 106353, 106496, 106656, 108108, 108878, 109824, 110292, 110592, 112112, 112896, 114048, 114534, 114688, 115544, 116352, 116424, 117936, 118272, 118776, 119808, 119988, 122304, 123552, 124432, 126048, 126126, 127008, 128128, 128674, 129024, 129987, 130896, 131072, 133056, 133623, 134784, 135168, 135744, 137592, 137984, 139776, 139986, 141804, 142208, 144144, 145152, 146432, 147056, 147456, 150528, 152064, 152712, 155136, 155232, 157248, 157696, 158368, 159744, 159984, 162162, 163072, 163317, 164736, 165438, 165888, 168064, 168168, 169344, 172032, 173316, 174528, 174636, 177408, 178164, 179712, 179982, 180224, 180992, 183456, 185328, 186368, 186648, 189072, 189189, 192192, 193011, 193536, 196608, 199584, 200704, 202202, 202752, 203616, 206388, 206848, 206976, 209664, 209979, 212706, 212992, 213312, 216216, 217756, 219648, 220584, 221184, 224224, 225792, 228096, 229068, 229376, 231088, 232704, 232848, 235872, 236544, 237552, 239616, 239976, 244608, 247104, 248157, 248864, 252096, 252252, 254016, 256256, 257348, 258048, 259974, 261792, 266112, 267246, 269568, 270336, 271488, 275184, 275968, 279552, 279972, 283608, 284416, 288288, 290304, 292864, 294112, 294912, 301056, 303303, 304128, 305424, 310272, 310464, 314496, 315392, 316736, 319488, 319968, 324324, 326144, 326634, 329472, 330876, 331776, 336128, 336336, 338688, 344064, 346632, 349056, 349272, 354816, 356328, 359424, 359964, 360448, 361984, 366912, 370656, 372736, 373296, 378144, 378378, 384384, 386022, 387072, 389961, 393216, 399168, 400869, 401408, 404404, 405504, 407232, 412776, 413696, 413952, 419328, 419958, 425412, 425984, 426624, 432432, 435512, 439296, 441168, 442368, 448448, 451584, 456192, 458136, 458752, 462176, 465408, 465696, 471744, 473088, 475104, 479232, 479952, 489216, 489951, 494208, 496314, 497728, 504192, 504504, 508032, 512512, 514696, 516096, 519948, 523584, 532224, 534492, 539136, 540672, 542976, 550368, 551936, 559104, 559944, 567216, 567567, 568832, 576576, 579033, 580608, 585728, 588224, 589824, 602112, 606606, 608256, 610848, 620544, 620928, 628992, 629937, 630784, 633472, 638976, 639936, 648648, 652288, 653268, 658944, 661752, 663552, 672256, 672672, 677376, 688128, 693264, 698112, 698544, 707707, 709632, 712656, 718848, 719928, 720896, 723968, 733824, 741312, 744471, 745472, 746592, 756288, 756756, 768768, 772044, 774144, 779922, 798336, 801738, 802816, 808808, 811008, 814464, 825552, 827392, 827904, 838656, 839916, 850824, 851968, 853248, 864864, 871024, 878592, 882336, 884736, 896896, 903168, 909909, 912384, 916272, 917504, 924352, 930816, 931392, 943488, 946176, 950208, 958464, 959904, 978432, 979902, 988416, 992628, 995456, 1008384, 1009008, 1016064, 1025024, 1029392, 1032192, 1039896, 1047168, 1064448, 1068984, 1078272, 1081344, 1085952, 1100736, 1103872, 1118208, 1119888, 1134432, 1135134, 1137664, 1153152, 1158066, 1161216, 1169883, 1171456, 1176448, 1179648, 1204224, 1213212, 1216512, 1221696, 1241088, 1241856, 1257984, 1259874, 1261568, 1266944, 1277952, 1279872, 1297296, 1304576, 1306536, 1317888, 1323504, 1327104, 1344512, 1345344, 1354752, 1376256, 1386528, 1396224, 1397088, 1415414, 1419264, 1425312, 1437696, 1439856, 1441792, 1447936, 1467648, 1469853, 1482624, 1488942, 1490944, 1493184, 1512576, 1513512, 1537536, 1544088, 1548288, 1559844, 1596672, 1603476, 1605632, 1617616, 1622016, 1628928, 1651104, 1654784, 1655808, 1677312, 1679832, 1701648, 1703936, 1706496, 1729728, 1737099, 1742048, 1757184, 1764672, 1769472, 1793792, 1806336, 1819818, 1824768, 1832544, 1848704, 1861632, 1862784, 1886976, 1892352, 1900416, 1916928, 1919808, 1956864, 1959804, 1976832, 1985256, 1990912, 2016768, 2018016, 2032128, 2050048, 2058784, 2064384, 2079792, 2094336, 2123121, 2128896, 2137968, 2156544, 2162688, 2171904, 2201472, 2207744, 2236416, 2239776, 2268864, 2270268, 2275328, 2306304, 2316132, 2322432, 2339766, 2342912, 2352896, 2408448, 2426424, 2433024, 2443392, 2482176, 2483712, 2515968, 2519748, 2523136, 2533888, 2555904, 2559744, 2594592, 2609152, 2613072, 2635776, 2647008, 2654208, 2689024, 2690688, 2709504, 2729727, 2752512, 2773056, 2792448, 2794176, 2830828, 2838528, 2850624, 2875392, 2879712, 2895872, 2935296, 2939706, 2965248, 2977884, 2981888, 2986368, 3025152, 3027024, 3075072, 3088176, 3096576, 3119688, 3193344, 3206952, 3211264, 3235232, 3244032, 3257856, 3302208, 3309568, 3311616, 3354624, 3359664, 3403296, 3412992, 3459456, 3474198, 3484096, 3514368, 3529344, 3538944, 3587584, 3612672, 3639636, 3649536, 3665088, 3697408, 3723264, 3725568, 3773952, 3784704, 3800832, 3833856, 3839616, 3913728, 3919608, 3953664, 3970512, 3981824, 4033536, 4036032, 4064256, 4100096, 4117568, 4128768, 4159584, 4188672, 4246242, 4257792, 4275936, 4313088, 4325376, 4343808, 4402944, 4409559, 4415488, 4472832, 4479552, 4537728, 4540536, 4550656, 4612608, 4632264, 4644864, 4679532, 4685824, 4705792, 4816896, 4852848, 4866048, 4886784, 4964352, 4967424, 5031936, 5039496, 5046272, 5067776, 5111808, 5119488, 5189184, 5211297, 5218304, 5226144, 5271552, 5294016, 5308416, 5378048, 5381376, 5419008, 5459454, 5546112, 5584896, 5588352, 5661656, 5677056, 5701248, 5750784, 5759424, 5791744, 5870592, 5879412, 5930496, 5955768, 5963776, 5972736, 6050304, 6054048, 6150144, 6176352, 6193152, 6239376, 6369363, 6386688, 6413904, 6422528, 6470464, 6488064, 6515712, 6604416, 6619136, 6623232, 6709248, 6719328, 6806592, 6825984, 6918912, 6948396, 6968192, 7028736, 7058688, 7175168, 7225344, 7279272, 7299072, 7330176, 7394816, 7446528, 7451136, 7547904, 7569408, 7601664, 7667712, 7679232, 7827456, 7839216, 7907328, 7941024, 7963648, 8067072, 8072064, 8128512, 8189181, 8200192, 8235136, 8257536, 8319168, 8377344, 8492484, 8515584, 8551872, 8626176, 8687616, 8805888, 8819118, 8830976, 8945664, 8959104, 9075456, 9081072, 9101312, 9225216, 9264528, 9289728, 9359064, 9371648, 9411584, 9633792, 9705696, 9732096, 9773568, 9928704, 9934848, 10063872, 10078992, 10092544, 10135552, 10238976, 10378368, 10422594, 10436608, 10452288, 10543104, 10588032, 10616832, 10756096, 10762752, 10838016, 10918908, 11092224, 11169792, 11176704, 11323312, 11354112, 11402496, 11501568, 11518848, 11583488, 11741184, 11758824, 11860992, 11911536, 11927552, 11945472, 12100608, 12108096, 12300288, 12352704, 12386304, 12478752, 12738726, 12773376, 12827808, 12940928, 12976128, 13031424, 13208832, 13238272, 13246464, 13418496, 13438656, 13613184, 13651968, 13837824, 13896792, 13936384, 14057472, 14117376, 14350336, 14450688, 14558544, 14598144, 14660352, 14789632, 14893056, 14902272, 15095808, 15138816, 15203328, 15335424, 15358464, 15654912, 15678432, 15814656, 15882048, 15927296, 16134144, 16144128, 16257024, 16378362, 16400384, 16470272, 16638336, 16754688, 16984968, 17031168, 17103744, 17252352, 17375232, 17611776, 17638236, 17661952, 17891328, 17918208, 18150912, 18162144, 18202624, 18450432, 18529056, 18579456, 18718128, 18743296, 18823168, 19108089, 19267584, 19411392, 19464192, 19547136, 19857408, 19869696, 20127744, 20157984, 20271104, 20477952, 20756736, 20845188, 20873216, 20904576, 21086208, 21176064, 21512192, 21525504, 21676032, 21837816, 22184448, 22339584, 22353408, 22646624, 22708224, 22804992, 23003136, 23037696, 23166976, 23482368, 23517648, 23721984, 23823072, 23890944, 24201216, 24216192, 24600576, 24705408, 24772608, 24957504, 25477452, 25546752, 25655616, 25881856, 26062848, 26417664, 26492928, 26836992, 26877312, 27226368, 27303936, 27675648, 27793584, 27872768, 28114944, 28234752, 28700672, 28901376, 29117088, 29196288, 29320704, 29579264, 29786112, 29804544, 30191616, 30277632, 30406656, 30716928, 31309824, 31356864, 31629312, 31764096, 31854592, 32268288, 32288256, 32514048, 32756724, 32800768, 32940544, 33276672, 33509376, 33969936, 34062336, 34207488, 34504704, 34750464, 35223552, 35276472, 35323904, 35782656, 35836416, 36301824, 36324288, 36405248, 36900864, 37058112, 37158912, 37436256, 37646336, 38216178, 38822784, 38928384, 39094272, 39714816, 39739392, 40255488, 40315968, 40542208, 40955904, 41513472, 41690376, 41746432, 41809152, 42172416, 42352128, 43024384, 43051008, 43352064, 43675632, 44368896, 44679168, 44706816, 45293248, 45416448, 45609984, 46006272, 46075392, 46333952, 46964736, 47035296, 47443968, 47646144, 47781888, 48402432, 48432384, 49201152, 49410816, 49915008, 50954904, 51093504, 51311232, 51763712, 52125696, 52835328, 52985856, 53673984, 53754624, 54452736, 54607872, 55351296, 55587168, 55745536, 56229888, 56469504, 57324267, 57401344, 57802752, 58234176, 58392576, 58641408, 59158528, 59572224, 59609088, 60383232, 60813312, 61433856, 62619648, 62713728, 63258624, 63528192, 63709184, 64536576, 64576512, 65028096, 65513448, 65601536, 65881088, 66553344, 67018752, 67939872, 68124672, 68414976, 69009408, 69500928, 70447104, 70552944, 70647808, 71672832, 72603648, 72648576, 72810496, 73801728, 74116224, 74317824, 74872512, 75292672, 76432356, 77645568, 78188544, 79478784, 80510976, 80631936, 81084416, 81911808, 83026944, 83380752, 83492864, 83618304, 84344832, 84704256, 86048768, 86102016, 86704128, 87351264, 88737792, 89358336, 89413632, 90586496, 90832896, 91219968, 92150784, 92667904, 93929472, 94070592, 94887936, 95292288, 95563776, 96804864, 96864768, 98402304, 98821632, 99830016, 101909808, 102187008, 102622464, 103527424, 104251392, 105670656, 105971712, 107347968, 107509248, 108905472, 109215744, 110702592, 111174336, 111491072, 112939008, 114648534, 114802688, 116468352, 116785152, 117282816, 118317056, 119144448, 119218176, 120766464, 121626624, 122867712, 125239296, 125427456, 126517248, 127056384, 127418368, 129073152, 129153024, 130056192, 131026896, 131203072, 131762176, 133106688, 134037504, 135879744, 136249344, 136829952, 138018816, 139001856, 140894208, 141105888, 143345664, 145207296, 145297152, 145620992, 147603456, 148232448, 149745024, 150585344, 152864712, 155291136, 156377088, 158957568, 161021952, 161263872, 162168832, 163823616, 166053888, 166761504, 167236608, 168689664, 169408512, 172097536, 172204032, 173408256, 174702528, 177475584, 178716672, 178827264, 181172992, 182439936, 184301568, 187858944, 188141184, 189775872, 190584576, 191127552, 193609728, 193729536, 196804608, 197643264, 199660032, 203819616, 204374016, 205244928, 207054848, 208502784, 211341312, 211943424, 215018496, 217810944, 218431488, 221405184, 222348672, 222982144, 225878016, 229297068, 229605376, 232936704, 234565632, 236634112, 238436352, 241532928, 243253248, 245735424, 250478592, 250854912, 253034496, 254112768, 254836736, 258146304, 258306048, 260112384, 262053792, 263524352, 266213376, 268075008, 271759488, 272498688, 273659904, 278003712, 281788416, 282211776, 286691328, 290414592, 290594304, 295206912, 296464896, 299490048, 301170688, 305729424, 310582272, 312754176, 317915136, 322043904, 322527744, 324337664, 327647232, 332107776, 333523008, 334473216, 338817024, 344408064, 349405056, 354951168, 357433344, 357654528, 362345984, 364879872, 368603136, 375717888, 376282368, 379551744, 381169152, 382255104, 387219456, 387459072, 393609216, 395286528, 399320064, 407639232, 408748032, 410489856, 414109696, 417005568, 422682624, 430036992, 435621888, 436862976, 442810368, 444697344, 445964288, 451756032, 458594136, 459210752, 465873408, 469131264, 473268224, 476872704, 483065856, 486506496, 491470848, 501709824, 506068992, 508225536, 509673472, 516292608, 516612096, 520224768, 524107584, 527048704, 532426752, 536150016, 543518976, 547319808, 563576832, 564423552, 573382656, 580829184, 581188608, 590413824, 592929792, 598980096, 602341376, 611458848, 621164544, 625508352, 635830272, 645055488, 648675328, 655294464, 664215552, 667046016, 668946432, 677634048, 688816128, 698810112, 709902336, 715309056, 724691968, 729759744, 737206272, 751435776, 752564736, 759103488, 762338304, 764510208, 774438912, 774918144, 790573056, 798640128, 815278464, 817496064, 820979712, 828219392, 834011136, 845365248, 860073984, 871243776, 885620736, 889394688, 891928576, 903512064, 917188272, 918421504, 931746816, 938262528, 946536448, 953745408, 966131712, 973012992, 982941696, 1003419648, 1016451072, 1019346944, 1033224192, 1048215168, 1054097408, 1064853504, 1072300032, 1087037952, 1094639616, 1127153664, 1128847104, 1146765312, 1161658368, 1162377216, 1180827648, 1185859584, 1197960192, 1204682752, 1222917696, 1242329088, 1251016704, 1290110976, 1310588928, 1328431104, 1334092032, 1337892864, 1355268096, 1377632256, 1397620224, 1419804672, 1430618112, 1449383936, 1459519488, 1474412544, 1505129472, 1518206976, 1524676608, 1529020416, 1548877824, 1549836288, 1581146112, 1597280256, 1630556928, 1641959424, 1656438784, 1690730496, 1720147968, 1742487552, 1771241472, 1778789376, 1783857152, 1807024128, 1834376544, 1863493632, 1876525056, 1893072896, 1907490816, 1946025984, 1965883392, 2006839296, 2032902144, 2066448384, 2096430336, 2108194816, 2129707008, 2174075904, 2189279232, 2254307328, 2257694208, 2293530624, 2323316736, 2324754432, 2371719168, 2395920384, 2445835392, 2484658176, 2502033408, 2580221952, 2656862208, 2668184064, 2675785728, 2710536192, 2755264512, 2795240448, 2839609344, 2861236224, 2898767872, 2919038976, 2948825088, 3010258944, 3049353216, 3058040832, 3099672576, 3162292224, 3194560512, 3261113856, 3283918848, 3312877568, 3381460992, 3440295936, 3484975104, 3542482944, 3557578752, 3567714304, 3614048256, 3668753088, 3726987264, 3753050112, 3931766784, 4013678592, 4065804288, 4132896768, 4192860672, 4216389632, 4259414016, 4348151808, 4378558464, 4515388416, 4587061248, 4646633472, 4649508864, 4743438336, 4791840768, 4891670784, 4969316352, 5160443904, 5313724416, 5336368128, 5351571456, 5421072384, 5590480896, 5679218688, 5722472448, 5797535744, 5838077952, 5897650176, 6020517888, 6098706432, 6199345152, 6324584448, 6389121024, 6522227712, 6567837696, 6625755136, 6762921984, 6880591872, 6969950208, 7115157504, 7135428608, 7337506176, 7453974528, 7506100224, 8027357184, 8131608576, 8265793536, 8385721344, 8432779264, 8518828032, 8696303616, 8757116928, 9030776832, 9174122496, 9299017728, 9486876672, 9583681536, 9783341568, 9938632704, 10320887808, 10627448832, 10672736256, 10703142912, 10842144768, 11180961792, 11595071488, 11795300352, 12041035776, 12197412864, 12398690304, 12649168896, 12778242048, 13044455424, 13135675392, 13251510272, 13761183744, 13939900416, 14230315008, 14675012352, 14907949056, 16054714368, 16263217152, 16771442688, 17037656064, 17392607232, 17514233856, 18061553664, 18598035456, 18973753344, 19167363072, 19566683136, 19877265408, 20641775616, 21345472512, 21406285824, 22361923584, 23190142976, 24082071552, 24394825728, 24797380608, 25298337792, 25556484096, 26088910848, 26271350784, 27522367488, 28460630016, 29350024704, 29815898112, 32109428736, 32526434304, 33542885376, 34785214464, 36123107328, 37196070912, 37947506688, 38334726144, 39133366272, 39754530816, 41283551232, 42690945024, 44723847168, 46380285952, 48164143104, 48789651456, 51112968192, 52177821696, 52542701568, 56921260032, 58700049408, 59631796224, 64218857472, 67085770752, 69570428928, 72246214656, 74392141824, 75895013376, 76669452288, 78266732544, 82567102464, 85381890048, 89447694336, 92760571904, 96328286208, 97579302912, 104355643392, 113842520064, 117400098816, 119263592448, 134171541504, 139140857856, 144492429312, 153338904576, 156533465088, 170763780096, 178895388672, 192656572416, 208711286784, 227685040128, 234800197632, 268343083008, 278281715712, 288984858624, 313066930176, 341527560192, 357790777344, 417422573568, 469600395264, 536686166016, 577969717248, 626133860352, 683055120384, 834845147136, 939200790528, 1073372332032, 1252267720704, 1878401581056, 2504535441408, 3756803162112, 7513606324224}},
	}

	for _, tc := range test_cases {
		divisors := Divisors(tc.n)
		AssertEqualArrays(t, divisors, tc.divisors)
	}
}
