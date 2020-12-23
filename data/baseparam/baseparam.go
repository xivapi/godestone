package baseparam

// BaseParam represents a stat value of a character
type BaseParam uint8

// FFXIV attribute.
const (
	None BaseParam = iota
	Strength
	Dexterity
	Vitality
	Intelligence
	Mind
	Piety
	HP
	MP
	TP
	GP
	CP
	PhysicalDamage
	MagicDamage
	Delay
	AdditionalEffect
	AttackSpeed
	BlockRate
	BlockStrength
	Tenacity
	AttackPower
	Defense
	DirectHitRate
	Evasion
	MagicDefense
	CriticalHitPower
	CriticalHitResilience
	CriticalHit
	CriticalHitEvasion
	SlashingResistance
	PiercingResistance
	BluntResistance
	ProjectileResistance
	AttackMagicPotency
	HealingMagicPotency
	EnhancementMagicPotency
	ElementalBonus
	FireResistance
	IceResistance
	WindResistance
	EarthResistance
	LightningResistance
	WaterResistance
	MagicResistance
	Determination
	SkillSpeed
	SpellSpeed
	Haste
	Morale
	Enmity
	EnmityReduction
	CarefulDesynthesis
	EXPBonus
	Regen
	Refresh
	MainAttribute
	SecondaryAttribute
	SlowResistance
	PetrificationResistance
	ParalysisResistance
	SilenceResistance
	BlindResistance
	PoisonResistance
	StunResistance
	SleepResistance
	BindResistance
	HeavyResistance
	DoomResistance
	ReducedDurabilityLoss
	IncreasedSpiritbondGain
	Craftsmanship
	Control
	Gathering
	Perception
)
