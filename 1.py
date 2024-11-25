import random

def jugar_piedra_papel_tijera():
  """
  Función para jugar piedra, papel o tijera contra la computadora.
  """

  opciones = ["piedra", "papel", "tijera"]
  usuario = input("Elige piedra, papel o tijera: ").lower()

  while usuario not in opciones:
    usuario = input("Opción inválida. Elige piedra, papel o tijera: ").lower()

  computadora = random.choice(opciones)

  print(f"La computadora eligió {computadora}.")

  if usuario == computadora:
    print("Empate!")
  elif (usuario == "piedra" and computadora == "tijera") or \
       (usuario == "papel" and computadora == "piedra") or \
       (usuario == "tijera" and computadora == "papel"):
    print("¡Ganaste!")
  else:
    print("¡Perdiste!")

if __name__ == "__main__":
  while True:
    jugar_piedra_papel_tijera()
    jugar_de_nuevo = input("¿Quieres jugar de nuevo? (s/n): ").lower()
    while jugar_de_nuevo not in ["s", "n"]:
      jugar_de_nuevo = input("Opción inválida. ¿Quieres jugar de nuevo? (s/n): ").lower()
    if jugar_de_nuevo == "n":
      break