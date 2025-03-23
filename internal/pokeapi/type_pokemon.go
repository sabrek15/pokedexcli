package pokeapi

type Pokemon struct {
	Abilities		[]struct {
		Ability	struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot	int `json:"slot"`
	}	`json:"abilities"`
	BaseExperience	int `json:"base_experience"`
	Cries			struct {
		Latest		string `json:"latest"`
		Legacy		string `json:"legacy"`
	}	`json:"cries"`
	Forms			[]struct {
		Name		string `json:"name"`
		URL			string `json:"url"`
	}	`json:"forms"`
	GameIndices		[]struct {
		GameIndex	int	`json:"game_index"`
		Version		struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		}	`json:"version"`
	}	`json:"game_indice"`
	Height			int `json:"height"`
	HeldItems		[]struct {
		Item	struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		}	`json:"item"`
		VersionDetails []struct {
			Rarity	int `json:"rarity"`
			Version	struct {
				Name	string `json:"name"`
				URL		string `json:"url"`
			}	`json:"version"`
		}	`json:"version_details"`
	}	`json:"held_items"`
	ID				int `json:"id"`
	IsDefault		bool `json:"is_default"`
	LocationAreaEncounters	string `json:"location_area_encounters"`
	Moves			[]struct {
		Move struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		}	`json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt	int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL	 string `json:"url"`
			}	`json:"move_learn_method"`
			Order 			int `json:"order"`
			VersionGroup	struct {
				Name string `json:"name"`
				URL	 string `json:"url"`
			}	`json:"version_group"`
		}	`json:"version_group_details"`
	}	`json:"moves"`
	Name			string `json:"name"`
	Order			int	`json:"order"`
	PastAbilites 	[]struct{}	`json:"past_abilities"`
	PastTypes		[]struct {
		Generation	struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		}	`json:"generation"`
		Types		[]struct {
			Slot	int `json:"solt"`
			Type	struct {
				Name	string `json:"name"`
				URL		string `json:"url"`
			}	`json:"type"`
		}	`json:"types"`
	}
	Species			struct {
		Name		string `json:"name"`
		URL			string `json:"url"`
	}	`json:"species"`
	Sprites			struct {
		BackDefault	string `json:"back_default"`
		BackFemale	string `json:"back_female"`
		BackShiny	string `json:"back_shiny"`
		BackShinyFemale	string `json:"back_shiny_female"`
		FrontDefault	string `json:"front_default"`
		FrontFemale	string `json:"front_female"`
		FrontShiny	string `json:"front_shiny"`
		FrontShinyFemale	string `json:"front_shiny_female"`
		Other		struct {
			DreamWorld	struct {
				FrontDefault	string `json:"front_default"`
				FrontFemale		string `json:"front_female"`
			}	`json:"dream_world"`
			Home		struct {
				FrontDefault	string `json:"front_default"`
				FrontFemale	string `json:"front_female"`
				FrontShiny	string `json:"front_shiny"`
				FrontShinyFemale	string `json:"front_shiny_female"`
			}	`json:"home"`
			OfficialArtwork	struct {
				FrontDefault	string `json:"front_default"`
				FrontShiny	string `json:"front_shiny"`
			}	`json:"official_artwork"`
			Showdown		struct {
				BackDefault	string `json:"back_default"`
				BackFemale	string `json:"back_female"`
				BackShiny	string `json:"back_shiny"`
				BackShinyFemale	string `json:"back_shiny_female"`
				FrontDefault	string `json:"front_default"`
				FrontFemale	string `json:"front_female"`
				FrontShiny	string `json:"front_shiny"`
				FrontShinyFemale	string `json:"front_shiny_female"`
			}	`json:"showdown"`
			Versions		struct {
				GenerationI	struct {
					RedBlue		struct {
						BackDefault		string `json:"back_default"`
						BackGrey		string `json:"back_grey"`
						BackTransparent		string `json:"back_transparent"`
						FrontDefault		string `json:"front_default"`
						FrontGrey		string `json:"front_grey"`
						FrontTransparent		string `json:"front_transparent"`
					}	`json:"red_blue"`
					Yellow		struct {
						BackDefault		string `json:"back_default"`
						BackGrey		string `json:"back_grey"`
						BackTransparent		string `json:"back_transparent"`
						FrontDefault		string `json:"front_default"`
						FrontGrey		string `json:"front_grey"`
						FrontTransparent		string `json:"front_transparent"`
					}	`json:"yellow"`
				}	`json:"generation_i"`
				GenerationIi struct {
					Crystal		struct {
						BackDefault		string `json:"back_default"`
						BackShiny		string `json:"back_shiny"`
						BackShinyTransparent		string `json:"back_shiny_transparent"`
						BackTransparent		string `json:"back_transparent"`
						FrontDefault 		string `json:"front_default"`
						FrontShiny		string `json:"front_shiny"`
						FrontShinyTransparent		string `json:"front_shiny_transparent"`
						FrontTransparent		string `json:"front_transparent"`
					}	`json:"crystal"`
					Gold		struct {
						BackDefault		string `json:"back_default"`
						BackShiny		string `json:"back_shiny"`
						FrontDefault 		string `json:"front_default"`
						FrontShiny		string `json:"front_shiny"`
						FrontTransparent		string `json:"front_transparent"`
					}	`json:"gold"`
					Sliver		struct {
						BackDefault		string `json:"back_default"`
						BackShiny		string `json:"back_shiny"`
						FrontDefault 		string `json:"front_default"`
						FrontShiny		string `json:"front_shiny"`
						FrontTransparent		string `json:"front_transparent"`
					}	`json:"sliver"`
				}	`json:"generation_ii"`
				GenerationIii struct {
					Emerald		struct {
						FrontDefault 		string `json:"front_default"`
						FrontShiny		string `json:"front_shiny"`
					}	`json:"emerald"`
					FireredLeafgreen		struct {
						BackDefault		string `json:"back_default"`
						BackShiny		string `json:"back_shiny"`
						FrontDefault 		string `json:"front_default"`
						FrontShiny		string `json:"front_shiny"`
					}	`json:"firered_leafgreen"`
					RubySapphire		struct {
						BackDefault		string `json:"back_default"`
						BackShiny		string `json:"back_shiny"`
						FrontDefault 		string `json:"front_default"`
						FrontShiny		string `json:"front_shiny"`
					}	`json:"ruby_sapphire"`
				}	`json:"generation_iii"`
				GenerationIv struct {
					DiamondPearl		struct {
						BackDefault	string `json:"back_default"`
						BackFemale	string `json:"back_female"`
						BackShiny	string `json:"back_shiny"`
						BackShinyFemale	string `json:"back_shiny_female"`
						FrontDefault	string `json:"front_default"`
						FrontFemale	string `json:"front_female"`
						FrontShiny	string `json:"front_shiny"`
						FrontShinyFemale	string `json:"front_shiny_female"`
					}	`json:"diamond_pearl"`
					HeartgoldSoulsliver		struct {
						BackDefault	string `json:"back_default"`
						BackFemale	string `json:"back_female"`
						BackShiny	string `json:"back_shiny"`
						BackShinyFemale	string `json:"back_shiny_female"`
						FrontDefault	string `json:"front_default"`
						FrontFemale	string `json:"front_female"`
						FrontShiny	string `json:"front_shiny"`
						FrontShinyFemale	string `json:"front_shiny_female"`
					}	`json:"heartgold_soulsliver"`
					Platinum		struct {
						BackDefault	string `json:"back_default"`
						BackFemale	string `json:"back_female"`
						BackShiny	string `json:"back_shiny"`
						BackShinyFemale	string `json:"back_shiny_female"`
						FrontDefault	string `json:"front_default"`
						FrontFemale	string `json:"front_female"`
						FrontShiny	string `json:"front_shiny"`
						FrontShinyFemale	string `json:"front_shiny_female"`
					}	`json:"platinum"`
				}	`json:"generation_iv"`
				GenerationV	struct {
					BlackWhite	struct {
						Animated	struct {
							BackDefault	string `json:"back_default"`
							BackFemale	string `json:"back_female"`
							BackShiny	string `json:"back_shiny"`
							BackShinyFemale	string `json:"back_shiny_female"`
							FrontDefault	string `json:"front_default"`
							FrontFemale	string `json:"front_female"`
							FrontShiny	string `json:"front_shiny"`
							FrontShinyFemale	string `json:"front_shiny_female"`
						}	`json:"animated"`
					}	`json:"black_white"`
						BackDefault	string `json:"back_default"`
						BackFemale	string `json:"back_female"`
						BackShiny	string `json:"back_shiny"`
						BackShinyFemale	string `json:"back_shiny_female"`
						FrontDefault	string `json:"front_default"`
						FrontFemale	string `json:"front_female"`
						FrontShiny	string `json:"front_shiny"`
						FrontShinyFemale	string `json:"front_shiny_female"`
				}	`json:"generation_v"`
				GenerationVi	struct {
					OmegarubyAlphasapphire	struct {
						FrontDefault	string `json:"front_default"`
						FrontFemale	string `json:"front_female"`
						FrontShiny	string `json:"front_shiny"`
						FrontShinyFemale	string `json:"front_shiny_female"`
					}	`json:"omegaruby_alphasapphire"`
					XY	struct {
						FrontDefault	string `json:"front_default"`
						FrontFemale	string `json:"front_female"`
						FrontShiny	string `json:"front_shiny"`
						FrontShinyFemale	string `json:"front_shiny_female"`
					}	`json:"x-y"`
				}	`json:"generation_vi"`
				GenerationVii	struct {
					Icons	struct {
						FrontDefault	string `json:"front_default"`
						FrontFemale	string `json:"front_female"`
					}	`json:"icons"`
					UltraSunUltraMoon	struct {
						FrontDefault	string `json:"front_default"`
						FrontFemale	string `json:"front_female"`
						FrontShiny	string `json:"front_shiny"`
						FrontShinyFemale	string `json:"front_shiny_female"`
					}	`json:"ultra-sun-ultra-moon"`
				}	`json:"generation_vii"`
				GenerationViii	struct {
					Icons	struct {
						FrontDefault	string `json:"front_default"`
						FrontFemale	string `json:"front_female"`
					}	`json:"icons"`
				}	`json:"generation_viii"`
			}	`json:"versions"`
		}	`json:"other"`
	}	`json:"sprites"`
	Stats		[]struct {
		BaseStat 	int `json:"base_stat"`
		Effort		int `json:"effort"`
		Stat		struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		}	`json:"stat"`
	}	`json:"stats"`
	Types		[]struct {
		Slot		int `json:"slot"`
		Type		struct {
			Name	string `json:"name"`
			URL		string `json:"url"`
		}	`json:"type"`
	}
	Weight		int `json:"weight"`
}