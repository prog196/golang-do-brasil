package main

import (
	"fmt"
		"math/rand"
			"time"
			)

			func main() {
				rand.Seed(time.Now().UnixNano())
					fmt.Println("=== CARA OU COROA ===")
						fmt.Println("1 - Cara")
							fmt.Println("2 - Coroa")
								fmt.Print("Escolha (1 ou 2): ")
									
										var escolha int
											fmt.Scanln(&escolha)
												
													resultado := rand.Intn(2) + 1
														
															opcoes := map[int]string{1: "Cara", 2: "Coroa"}
																
																	fmt.Printf("Você escolheu: %s\n", opcoes[escolha])
																		fmt.Printf("Resultado: %s\n", opcoes[resultado])
																			
																				if escolha == resultado {
																						fmt.Println("VOCÊ GANHOU!")
																							} else {
																									fmt.Println("VOCÊ PERDEU!")
																										}
																										}