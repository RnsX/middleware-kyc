package dto

type TyronCheck struct {
	Name    string
	Surname string
	DOB     string
	Address string
}

type TyronCheckResult struct {
	HitNoHit int // 1 hit, 0 no hit
}
