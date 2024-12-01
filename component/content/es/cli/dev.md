# otter desarrollador

Esto ejecutará un servidor de desarrollo, el servidor de desarrollo es crucial para un desarrollo fluido debido a su recarga en caliente.

A medida que desarrollamos la interfaz de usuario e iteramos innumerables veces por minuto, necesitamos administrar la recarga en caliente para que podamos concentrarnos en escribir la interfaz de usuario y no distraernos con cosas inútiles.

```
otter dev
```

`otter dev` leerá su otter.json, sus variables de entorno y su .env para configurar un servidor de desarrollo.

El puerto se leerá desde una variable de entorno PORT o, alternativamente, puede usar un indicador `--port`
```
otter dev --port=8123
```
