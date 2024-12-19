import random

reglas = {
    "piedra": {"tijera": "ganaste", "papel": "perdiste"},
    "papel": {"piedra": "ganaste", "tijera": "perdiste"},
    "tijera": {"papel": "ganaste", "piedra": "perdiste"}
}

def obtener_eleccion_usuario():
    """Obtiene la elección del usuario con validación de entrada."""
    opciones = ["piedra", "papel", "tijera"]
    while True:
        try:
            eleccion = input("Elige piedra, papel o tijera: ").lower()
            if eleccion in opciones:
                return eleccion
            else:
                print("Opción inválida. Intenta de nuevo.")
        except ValueError:
            print("Entrada inválida. Intenta de nuevo.")

def obtener_eleccion_computadora():
    """Genera una elección aleatoria para la computadora."""
    opciones = ["piedra", "papel", "tijera"]
    return random.choice(opciones)

def determinar_ganador(usuario, computadora):
    """Determina el ganador de la ronda."""
    resultado = reglas.get(usuario, {}).get(computadora)
    if resultado:
        return resultado.capitalize()
    elif usuario == computadora:
        return "Empate"

def jugar_ronda():
    """Juega una ronda de piedra, papel o tijera."""
    usuario = obtener_eleccion_usuario()
    computadora = obtener_eleccion_computadora()
    print(f"La computadora eligió {computadora}.")
    resultado = determinar_ganador(usuario, computadora)
    print(f"¡{resultado}!")

if __name__ == "__main__":
    while True:
        jugar_ronda()
        jugar_de_nuevo = input("¿Quieres jugar de nuevo? (s/n): ").lower()
        while jugar_de_nuevo not in ["s", "n"]:
            jugar_de_nuevo = input("Opción inválida. ¿Quieres jugar de nuevo? (s/n): ").lower()
        if jugar_de_nuevo == "n":
            print("¡Gracias por jugar!")
            break
          
