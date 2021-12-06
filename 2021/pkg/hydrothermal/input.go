package hydrothermal

var inputs = []string{
	"976,35 -> 24,987",
	"552,172 -> 870,490",
	"647,640 -> 841,834",
	"580,460 -> 580,749",
	"614,575 -> 746,575",
	"97,846 -> 441,846",
	"467,680 -> 767,680",
	"722,860 -> 722,98",
	"31,338 -> 31,581",
	"113,712 -> 184,712",
	"738,897 -> 136,897",
	"820,750 -> 144,74",
	"291,411 -> 641,411",
	"581,878 -> 581,657",
	"449,540 -> 787,202",
	"79,925 -> 981,23",
	"800,120 -> 36,884",
	"253,603 -> 253,643",
	"574,138 -> 574,966",
	"847,199 -> 144,902",
	"816,177 -> 243,750",
	"963,632 -> 472,141",
	"38,41 -> 986,989",
	"980,225 -> 980,801",
	"255,350 -> 647,350",
	"732,311 -> 732,907",
	"109,662 -> 113,662",
	"333,317 -> 470,180",
	"111,146 -> 339,146",
	"136,856 -> 534,458",
	"555,39 -> 555,895",
	"699,327 -> 699,496",
	"280,948 -> 660,948",
	"919,293 -> 316,896",
	"343,645 -> 620,368",
	"14,984 -> 975,23",
	"219,653 -> 696,176",
	"50,350 -> 50,956",
	"919,550 -> 919,568",
	"405,532 -> 238,532",
	"328,95 -> 979,746",
	"564,716 -> 119,716",
	"52,285 -> 52,126",
	"240,671 -> 963,671",
	"691,416 -> 676,431",
	"216,247 -> 216,530",
	"103,309 -> 103,643",
	"265,163 -> 981,879",
	"623,399 -> 760,262",
	"392,568 -> 674,286",
	"280,82 -> 863,665",
	"657,522 -> 657,858",
	"194,16 -> 443,16",
	"158,326 -> 158,372",
	"582,530 -> 582,159",
	"857,638 -> 857,807",
	"463,575 -> 463,108",
	"74,390 -> 74,967",
	"437,892 -> 224,892",
	"854,409 -> 366,897",
	"875,858 -> 875,871",
	"245,683 -> 735,193",
	"911,285 -> 216,980",
	"944,170 -> 701,170",
	"149,244 -> 149,653",
	"879,926 -> 70,117",
	"193,198 -> 777,782",
	"11,581 -> 287,305",
	"513,163 -> 939,163",
	"551,705 -> 551,636",
	"546,79 -> 546,630",
	"392,877 -> 392,240",
	"859,821 -> 859,975",
	"755,581 -> 755,722",
	"941,636 -> 976,636",
	"212,807 -> 595,807",
	"24,847 -> 24,248",
	"972,46 -> 40,978",
	"434,328 -> 491,328",
	"573,663 -> 16,663",
	"882,43 -> 882,777",
	"162,786 -> 11,786",
	"774,340 -> 322,340",
	"33,775 -> 883,775",
	"422,958 -> 212,748",
	"973,879 -> 415,321",
	"278,602 -> 435,759",
	"983,99 -> 321,99",
	"594,502 -> 727,635",
	"755,674 -> 314,233",
	"67,336 -> 702,336",
	"617,128 -> 617,287",
	"735,929 -> 165,929",
	"758,778 -> 758,679",
	"38,27 -> 971,960",
	"873,419 -> 949,419",
	"305,138 -> 978,811",
	"282,404 -> 377,404",
	"753,267 -> 945,267",
	"877,796 -> 64,796",
	"35,32 -> 949,946",
	"18,383 -> 64,429",
	"855,407 -> 938,324",
	"845,965 -> 88,208",
	"77,960 -> 960,77",
	"225,714 -> 490,714",
	"619,40 -> 395,40",
	"87,379 -> 87,178",
	"961,828 -> 302,828",
	"78,321 -> 78,816",
	"243,620 -> 883,620",
	"581,560 -> 69,560",
	"420,957 -> 768,957",
	"927,427 -> 908,408",
	"100,406 -> 100,736",
	"369,27 -> 199,27",
	"177,804 -> 177,727",
	"83,807 -> 166,724",
	"358,119 -> 358,583",
	"866,223 -> 348,741",
	"283,636 -> 283,476",
	"792,481 -> 161,481",
	"126,476 -> 612,962",
	"829,437 -> 829,444",
	"402,683 -> 402,11",
	"680,278 -> 676,278",
	"391,597 -> 521,467",
	"467,787 -> 646,608",
	"637,689 -> 637,959",
	"161,173 -> 161,604",
	"582,252 -> 582,181",
	"971,93 -> 329,93",
	"763,195 -> 156,802",
	"576,504 -> 755,325",
	"156,56 -> 657,557",
	"276,940 -> 836,380",
	"800,933 -> 800,734",
	"486,607 -> 486,54",
	"847,679 -> 299,131",
	"558,711 -> 558,643",
	"44,869 -> 44,877",
	"897,399 -> 897,265",
	"856,217 -> 856,701",
	"395,784 -> 395,634",
	"443,647 -> 443,977",
	"59,735 -> 59,860",
	"564,519 -> 173,910",
	"516,860 -> 54,860",
	"23,467 -> 23,551",
	"82,102 -> 849,869",
	"316,551 -> 195,551",
	"943,41 -> 25,959",
	"314,865 -> 314,74",
	"434,491 -> 501,491",
	"941,563 -> 860,563",
	"937,842 -> 320,225",
	"415,725 -> 415,841",
	"822,308 -> 500,308",
	"136,434 -> 22,434",
	"275,356 -> 280,356",
	"672,935 -> 22,935",
	"776,22 -> 55,743",
	"219,198 -> 219,775",
	"977,923 -> 977,344",
	"37,922 -> 815,144",
	"107,493 -> 107,804",
	"840,913 -> 840,686",
	"249,774 -> 249,485",
	"765,696 -> 649,696",
	"491,708 -> 302,708",
	"345,589 -> 345,357",
	"935,206 -> 759,206",
	"757,823 -> 30,96",
	"590,513 -> 746,513",
	"373,18 -> 845,490",
	"816,829 -> 816,608",
	"104,15 -> 964,875",
	"650,675 -> 650,389",
	"333,987 -> 857,463",
	"627,398 -> 627,223",
	"578,208 -> 159,208",
	"355,594 -> 355,633",
	"63,921 -> 873,921",
	"510,434 -> 801,143",
	"953,928 -> 308,283",
	"947,48 -> 25,970",
	"384,203 -> 384,97",
	"806,160 -> 934,288",
	"690,29 -> 269,29",
	"825,834 -> 269,278",
	"620,613 -> 620,540",
	"554,698 -> 419,833",
	"887,554 -> 457,554",
	"276,573 -> 276,487",
	"213,211 -> 213,619",
	"437,621 -> 141,917",
	"951,59 -> 951,272",
	"270,455 -> 270,336",
	"727,25 -> 42,710",
	"803,384 -> 615,196",
	"643,715 -> 643,741",
	"750,815 -> 642,923",
	"464,693 -> 714,943",
	"828,773 -> 189,134",
	"507,858 -> 58,858",
	"289,898 -> 190,898",
	"380,518 -> 749,149",
	"696,219 -> 760,219",
	"678,177 -> 686,185",
	"241,103 -> 857,103",
	"782,773 -> 782,508",
	"20,24 -> 277,281",
	"175,805 -> 59,805",
	"375,944 -> 375,938",
	"180,971 -> 203,971",
	"379,984 -> 830,984",
	"298,376 -> 254,376",
	"807,376 -> 486,376",
	"931,512 -> 931,931",
	"889,859 -> 361,859",
	"632,546 -> 298,880",
	"429,616 -> 583,770",
	"814,838 -> 503,527",
	"64,301 -> 753,301",
	"706,124 -> 706,698",
	"323,976 -> 323,43",
	"42,82 -> 550,590",
	"260,528 -> 260,462",
	"201,656 -> 593,656",
	"348,516 -> 203,516",
	"201,675 -> 413,675",
	"928,70 -> 138,860",
	"323,427 -> 601,427",
	"874,156 -> 630,156",
	"335,374 -> 335,522",
	"237,551 -> 597,551",
	"14,125 -> 909,125",
	"805,59 -> 67,797",
	"656,684 -> 656,263",
	"487,544 -> 487,464",
	"637,890 -> 637,606",
	"27,983 -> 952,58",
	"899,93 -> 77,915",
	"504,288 -> 504,689",
	"404,289 -> 700,289",
	"643,336 -> 321,336",
	"190,865 -> 674,865",
	"844,12 -> 81,775",
	"821,365 -> 821,453",
	"503,20 -> 503,811",
	"20,122 -> 983,122",
	"28,231 -> 398,231",
	"441,263 -> 931,263",
	"130,19 -> 925,19",
	"577,873 -> 577,706",
	"322,489 -> 322,621",
	"269,134 -> 935,800",
	"61,841 -> 491,841",
	"286,720 -> 542,464",
	"497,530 -> 497,266",
	"178,616 -> 512,282",
	"184,338 -> 184,241",
	"906,946 -> 327,946",
	"879,947 -> 879,302",
	"815,788 -> 963,788",
	"791,322 -> 791,395",
	"851,116 -> 793,116",
	"232,114 -> 934,816",
	"273,839 -> 157,839",
	"184,876 -> 184,138",
	"298,586 -> 634,250",
	"130,127 -> 130,753",
	"453,485 -> 855,887",
	"663,776 -> 934,776",
	"799,326 -> 799,661",
	"56,498 -> 274,716",
	"650,317 -> 52,915",
	"93,342 -> 391,44",
	"972,22 -> 59,935",
	"761,78 -> 508,331",
	"577,578 -> 15,16",
	"51,902 -> 72,881",
	"51,91 -> 51,422",
	"89,602 -> 89,280",
	"339,129 -> 339,329",
	"173,413 -> 489,413",
	"756,383 -> 745,383",
	"216,39 -> 216,373",
	"844,404 -> 552,404",
	"313,276 -> 313,895",
	"236,330 -> 231,330",
	"836,496 -> 836,291",
	"18,48 -> 950,980",
	"562,408 -> 562,606",
	"70,105 -> 70,469",
	"402,458 -> 694,166",
	"228,644 -> 689,183",
	"220,646 -> 834,32",
	"691,734 -> 141,184",
	"24,974 -> 978,20",
	"805,111 -> 11,905",
	"765,765 -> 210,210",
	"265,810 -> 248,810",
	"742,506 -> 179,506",
	"945,954 -> 40,49",
	"403,464 -> 600,464",
	"149,784 -> 754,784",
	"784,757 -> 222,757",
	"905,839 -> 160,839",
	"660,971 -> 609,971",
	"148,505 -> 309,505",
	"571,494 -> 323,494",
	"573,109 -> 71,109",
	"513,649 -> 54,649",
	"287,582 -> 287,604",
	"569,218 -> 569,790",
	"108,740 -> 108,816",
	"542,899 -> 445,802",
	"939,519 -> 939,752",
	"810,643 -> 810,236",
	"571,95 -> 560,106",
	"547,431 -> 547,42",
	"299,536 -> 299,820",
	"506,808 -> 437,808",
	"727,803 -> 941,589",
	"583,100 -> 891,100",
	"552,556 -> 515,556",
	"789,871 -> 90,172",
	"81,649 -> 641,89",
	"264,518 -> 511,518",
	"498,324 -> 596,324",
	"642,836 -> 642,481",
	"695,827 -> 768,900",
	"603,940 -> 603,352",
	"975,64 -> 92,947",
	"65,515 -> 65,405",
	"766,667 -> 344,667",
	"284,162 -> 245,162",
	"139,123 -> 942,926",
	"316,906 -> 316,907",
	"42,418 -> 224,600",
	"338,733 -> 338,46",
	"448,744 -> 448,796",
	"198,153 -> 198,723",
	"122,433 -> 712,433",
	"887,708 -> 685,708",
	"452,265 -> 817,630",
	"317,613 -> 317,959",
	"185,841 -> 788,238",
	"702,558 -> 734,558",
	"45,749 -> 330,464",
	"250,174 -> 250,561",
	"276,664 -> 793,664",
	"164,434 -> 619,434",
	"360,13 -> 686,339",
	"52,333 -> 361,642",
	"315,675 -> 315,175",
	"646,530 -> 815,699",
	"363,554 -> 58,554",
	"730,80 -> 38,772",
	"90,85 -> 494,85",
	"863,64 -> 218,64",
	"633,492 -> 633,134",
	"321,919 -> 324,919",
	"395,133 -> 395,592",
	"152,963 -> 983,132",
	"42,149 -> 674,781",
	"754,146 -> 858,146",
	"53,628 -> 976,628",
	"433,365 -> 433,735",
	"951,360 -> 951,913",
	"875,250 -> 875,463",
	"923,348 -> 208,348",
	"951,586 -> 103,586",
	"818,924 -> 178,284",
	"265,130 -> 265,859",
	"26,410 -> 805,410",
	"847,149 -> 190,806",
	"136,36 -> 797,36",
	"841,660 -> 620,660",
	"759,553 -> 393,919",
	"530,743 -> 647,860",
	"163,909 -> 979,93",
	"798,175 -> 399,574",
	"934,847 -> 934,231",
	"373,749 -> 373,397",
	"679,871 -> 695,887",
	"407,468 -> 524,468",
	"890,611 -> 988,611",
	"104,706 -> 745,65",
	"533,659 -> 533,126",
	"342,460 -> 187,460",
	"398,26 -> 398,254",
	"116,11 -> 886,781",
	"846,317 -> 329,834",
	"919,104 -> 120,903",
	"93,843 -> 912,24",
	"618,610 -> 618,311",
	"834,276 -> 85,276",
	"983,26 -> 42,967",
	"412,706 -> 412,204",
	"51,966 -> 952,65",
	"969,871 -> 969,258",
	"51,652 -> 608,95",
	"289,903 -> 869,903",
	"283,760 -> 781,760",
	"521,74 -> 521,848",
	"720,572 -> 638,572",
	"146,847 -> 146,930",
	"980,953 -> 42,15",
	"49,956 -> 941,64",
	"209,242 -> 905,242",
	"249,185 -> 374,60",
	"916,738 -> 916,793",
	"12,922 -> 872,62",
	"543,198 -> 123,198",
	"316,423 -> 316,549",
	"694,514 -> 869,514",
	"36,46 -> 916,926",
	"427,295 -> 267,295",
	"884,669 -> 884,175",
	"558,379 -> 549,379",
	"89,85 -> 968,964",
	"48,544 -> 48,675",
	"337,81 -> 337,870",
	"953,46 -> 44,955",
	"37,912 -> 251,698",
	"342,191 -> 342,238",
	"874,289 -> 353,810",
	"547,935 -> 97,485",
	"392,359 -> 859,826",
	"329,815 -> 329,98",
	"65,22 -> 118,75",
	"803,341 -> 803,763",
	"389,98 -> 964,98",
	"420,520 -> 396,520",
	"204,650 -> 583,650",
	"446,77 -> 739,77",
	"208,447 -> 75,580",
	"693,443 -> 693,668",
	"341,697 -> 802,697",
	"398,718 -> 318,638",
	"430,38 -> 430,665",
	"519,932 -> 490,932",
	"381,492 -> 242,353",
	"896,616 -> 638,616",
	"520,552 -> 901,933",
	"750,44 -> 307,44",
	"653,209 -> 186,676",
	"399,447 -> 208,256",
	"741,146 -> 741,399",
	"228,893 -> 228,214",
	"934,925 -> 71,62",
	"986,74 -> 74,986",
	"422,88 -> 467,43",
	"566,680 -> 875,371",
	"328,465 -> 38,465",
	"705,620 -> 705,441",
	"534,256 -> 534,784",
	"909,939 -> 909,222",
	"467,640 -> 323,640",
	"372,725 -> 816,281",
	"78,631 -> 78,933",
	"739,376 -> 397,718",
	"901,954 -> 153,206",
	"869,212 -> 799,212",
	"192,946 -> 192,736",
	"946,13 -> 13,946",
	"267,480 -> 267,561",
	"954,287 -> 954,190",
	"145,935 -> 913,167",
	"295,152 -> 458,152",
	"10,690 -> 139,690",
	"121,23 -> 977,879",
	"265,247 -> 265,976",
	"281,793 -> 787,793",
	"988,355 -> 367,976",
	"97,807 -> 323,807",
	"527,506 -> 474,506",
	"359,340 -> 359,280",
	"371,203 -> 801,203",
	"53,593 -> 53,980",
	"377,705 -> 987,95",
	"901,975 -> 153,227",
	"851,442 -> 565,442",
	"425,976 -> 850,551",
	"766,674 -> 766,572",
	"18,757 -> 18,444",
	"386,682 -> 386,424",
	"966,640 -> 604,278",
	"919,973 -> 141,195",
	"672,768 -> 405,768",
	"271,814 -> 971,114",
	"719,902 -> 474,902",
	"365,768 -> 877,256",
	"360,787 -> 214,787",
	"133,616 -> 266,483",
	"577,399 -> 59,399",
	"290,74 -> 290,145",
	"154,131 -> 154,210",
}
