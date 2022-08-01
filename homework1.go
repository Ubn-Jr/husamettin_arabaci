package main

import "fmt"

//TODO
//Pull Request Düzenlemesi
//DONE

func main() {

	//Toplam alanı tutmak için int türünde bir değişken tanımlanıyor
	//--var alan int
	var area int

	//Bodrum katında bir alanın varlığını kontrol etmek için bool türünde bir değişken tanımlanıyor
	//--var bodrumVarMi bool
	var isBase bool

	//Binamızın kat sayısını tutmak için int türünde bir değişken tanımlanıyor
	//--katSayisi:=7
	fCount := 7

	//Apartmanımızın 4 cephesindede daire olup olmadığını kontrol etmek için bool türünde bir değişken tanımlanıyor
	// apartman4CepheliMi := true
	isApt4Site := true

	//Apartmanımızın en üst katındaki dairelerin dublex olup olmadığnı kontrol etmek için bool türünde bir değişken tanımlanıyor
	//--doublexVarMi := true
	isDoublexEx := true

	//Bodrum katında bir alanın varlığı ayarlanıyor
	//--bodrumVarMi = true
	isBase = true

	//Sitemizdeki apartman sayısını tutmak için int türünde bir değişken tanımlanıyor
	//--apartmanSayisi := 5
	aptCount := 5

	//Sitemizdeki bir binanın toplam m2 olarak yaşam alanı hesaplanıyor
	//--alan = toplamAlanHesapla(katSayisi,apartman4CepheliMi,doublexVarMi,bodrumVarMi)
	area = totalAreaCalc(fCount, isApt4Site, isDoublexEx, isBase)

	//Bir binanın toplam yaşam alanı ekrana yazdırılıyor
	//--fmt.Println("Bir Apartmanın Yaşam Alanı : ", alan)
	fmt.Println("Area By Apt : ", area)

	//Sitemizdeki tüm binaların toplam yaşam alanını tutmak için int türünde bir değişken tanımlanıyor
	//ve tanımlanan değişkene tüm binaların toplam yaşam alanı kaydediliyor
	//--siteToplamAlani :=  siteAlaniHesapla(apartmanSayisi,alan)
	siteTotalArea := siteAreaCalc(aptCount, area)

	//Sitemizdeki toplam yaşam alanı ekrana yazdırılıyor
	//--fmt.Println("Site Toplam Alanı : ", siteToplamAlani)
	fmt.Println("Area By Site : ", siteTotalArea)

}

//Verilen apartman sayisi ve apartman başı yaşam alanı değerine göre toplam site alanını hesaplayan fonksiyon
//--func siteAlaniHesapla(apartmanSayisi int, apartmanAlani int) int {
func siteAreaCalc(aptCount int, aptArea int) int {

	// İşlem sonucu olarak geriye apartman sayısı ile apartman yaşam alanının çarpımını döndürüyor
	//--return apartmanSayisi * apartmanAlani
	return aptCount * aptArea
}

//Verilen apartman katsayısı , apartmanın 4 cephesindede daire olup olmadığının bilgisi
// apartmanda dublex kat olup olmadığının bilgisi ve bodrum kat olup olmadığının bilgisi ile
// bir apartmanın toplam yaşam alanını hesaplayan fonksiyon
//--func toplamAlanHesapla(katSayisi int, apartman4CepheliMi bool, doublexVarMi bool, bodrumVarMi bool) int {
func totalAreaCalc(floorCount int, isApt4Site bool, isExDoublex bool, isBasemend bool) int {

	//Hespalama sonucunda fonksiyon çıktısı olarak kullanıcak alanı tutmak için int türünde bir değişken tanımlanıyor
	//--toplamAlan :=0
	totalArea := 0

	//Normal bir dairenin verilen özelliklere göre yaşam alanının tutmak için int türünde bir değişken tanımlanıyor
	//Ve tanımlanan değişkene normal bir dairenin yaşam alanı kaydediliyor
	//--normalDaireM2 := getirDaireM2(apartman4CepheliMi,false)
	normalSuiteM2 := getSuiteM2(isApt4Site, false)

	//Apartmanın cephe durumuna göre normal katlardaki bir katta kaç daire olduğunun değerini tutmak için
	// int türünde bir değişken tanımlanıyor
	//Ve tanımlanan değişkene normal katlardaki bir katta kaç daire olduğunun değeri kaydediliyor
	//--katBasiNormalDaireSayisi := getirKatBasiDaireSayisi(apartman4CepheliMi,true)
	normalSuiteCount := getSuiteCount(isApt4Site, true)

	//Bina alanı hesaplarken binada dublex katının olup olmadığı kontrol ediliyor
	//--if doublexVarMi {
	if isExDoublex {
		// Eğer binada dubleks kat varsa bu kodlar çalışır

		//Normal dairelerin toplam alanını tutmak için int türünde bir değişken tanımlanıyor
		// Ve kat sayısı 1 azaltılarak (dubleks katı sayılmayarak) kattaki daire sayısı ve dairelerin m2 değerleri çarpılıyor
		// çıkan sonuc değişkene kaydediliyor
		//--normalDairelerinToplamAlani = (katSayisi -1) * katBasiNormalDaireSayisi * normalDaireM2
		normalSuiteTotalArea := (floorCount - 1) * normalSuiteCount * normalSuiteM2

		//Dubleks bir dairenin verilen özelliklere göre yaşam alanının tutmak için int türünde bir değişken tanımlanıyor
		//Ve tanımlanan değişkene dubleks bir dairenin yaşam alanı kaydediliyor
		//--dubleksDaireM2 := getirDaireM2(apartman4CepheliMi,true)
		doublexSuiteM2 := getSuiteM2(isApt4Site, true)

		//Dubleks dairelerin toplam alanını tutmak için int türünde bir değişken tanımlanıyor
		// Ve kattaki dubleks daire sayısı ve dubleks dairelerin m2 değerleri çarpılıyor
		// çıkan sonuc değişkene kaydediliyor
		//--dubleksDairelerinToplamAlani = getirDaireM2(apartman4CepheliMi, true) * dubleksDaireM2
		dublekSuiteTotalArea := getSuiteCount(isApt4Site, true) * doublexSuiteM2

		//Toplam alan değişkenin normal dairelerin ve dubleks dairelerin toplam alanı kaydeiliyor
		//--toplamAlan = normalDairelerinToplamAlani + dubleksDairelerinToplamAlani
		totalArea = normalSuiteTotalArea + dublekSuiteTotalArea
	} else {
		// Eğer binada dublek kat yoksa bu kodlar çalışır

		//Toplam alan değişkenine normal dairelerin toplam alanı kaydeiliyor
		//--toplamAlan = katSayisi * katBasiNormalDaireSayisi * normalDaireM2
		totalArea = floorCount * normalSuiteCount * normalSuiteM2
	}

	//Bina alanı hesaplarken binada bodrum katının olup olmadığı kontrol ediliyor
	//--if bodrumVarMi {
	if isBasemend {
		// Eğer binada bodrum kat varsa bu kodlar çalışır

		//--toplamAlan += 70
		totalArea += 70
	}

	//Ve geriye toplam değeri döndürülüyor
	//--return toplamAlan
	return totalArea
}

// Verilen parametreler göre dairenin yaşam alanını geri döndüren fonksiyon
//--getirDaireM2(apartman4CepheliMi bool, dubleksDaireMi bool)
func getSuiteM2(isApt4Site bool, isDoublex bool) int {

	//Geri döndürülecek yaşam alanını tutmak için int türünde bir değişken tanımlanıyor
	var m2 int

	// Sorgulanan dairenin dubleks olup olmadığına bakılıyor
	//--if dubleksDaireMi {
	if isDoublex {
		//Sorgulanan daire dubleks ise bu kod çalışır

		//Apartmanın kaç cephesinde daire olduğuna bakılıyor
		//--if apartman4CepheliMi {
		if isApt4Site {
			// Eğer apartmanın 4 cephesinde daire varsa bu kod çalışır

			//m2 değerine 4 cepheli bir apartmanın dubleks dairesisnin yaşan alanı kaydediliyor
			m2 = 220
		} else {
			// Eğer apartmanın 2 cephesinde daire varsa bu kod çalışır

			//m2 değerine 2 cepheli bir apartmanın dubleks dairesisnin yaşan alanı kaydediliyor
			m2 = 350
		}
	} else {
		//Sorgulanan daire normal ise bu kod çalışır

		//Apartmanın kaç cephesinde daire olduğuna bakılıyor
		//--if apartman4CepheliMi {
		if isApt4Site {
			// Eğer apartmanın 4 cephesinde daire varsa bu kod çalışır

			//m2 değerine 4 cepheli bir apartmanın normal dairesinin yaşam alanı kaydediliyor
			m2 = 120
		} else {
			// Eğer apartmanın 2 cephesinde daire varsa bu kod çalışır

			//m2 değerine 2 cepheli bir apartmanın normal dairesinin yaşam alanı kaydediliyor
			m2 = 200
		}
	}

	//Kaydedilen yaşam alanı geriye döndürülüyor
	return m2
}

// Verilen parametreler göre kattaki daire sayısını geri döndüren fonksiyon
//--getirKatBasiDaireSayisi(apartman4CepheliMi bool, normalDaireMi bool)
func getSuiteCount(isApt4Site bool, isNormal bool) int {

	//Apartmanın 4 cephelimi ve normal katlar için mi kat başı daire sayısının istendiğine bakılıyor
	//--if apartman4CepheliMi && normalDaireMi
	if isApt4Site && isNormal {

		//Eğer öğrenilmek istenen kat başı daire sayısı
		// 4 cepheli bir apartmanın normal dairelerinin olduğu bir kat ise bu kod çalışır

		return 4
	} else {

		//Eğer 2 cepheli bir apartmanın normal katlarındaki veya her hangi bir binanın dubleks
		// katındaki dairelerin sayısı öğrenilmek isteniyorsa bu kodlar çalışır

		return 2
	}
}
