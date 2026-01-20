# Integración TPV Virtual Redsys (Entorno de Pruebas)

## 1. Configuración inicial

1.1. Edita el archivo `config/redsys.json` con los datos de pruebas de Redsys:
```json
{
  "merchant_code": "999008881",
  "terminal": "1",
  "secret_key": "qwertyasdf0123456789qwertyasdf0123456789qwertyasdf0123",
  "currency": "978",
  "test_url": "https://sis-t.redsys.es:25443/sis/realizarPago"
}
```

## 2. Endpoints disponibles

- **Iniciar pago:**
  - `GET /redsys/pay`
  - Genera y envía automáticamente el formulario de pago a Redsys (simula un pago de 1,00€).

- **Notificación de Redsys (IPN):**
  - `POST /redsys/notify`
  - Endpoint donde Redsys enviará la notificación del resultado del pago.

## 3. Flujo de pago desde el frontend

1. El usuario pulsa un botón "Pagar" en tu frontend.
2. El frontend hace una petición `GET` a `/redsys/pay` (puede ser un enlace o redirección).
3. El backend responde con un formulario HTML que se autoenvía a Redsys.
4. El usuario realiza el pago en la pasarela de Redsys.
5. Redsys redirige al usuario a `/redsys/ok` o `/redsys/ko` según el resultado (puedes implementar estos endpoints para mostrar mensajes personalizados).
6. Redsys envía una notificación `POST` a `/redsys/notify` con los datos del pago (debes procesar y validar la firma en este endpoint).

## 4. Información necesaria desde el frontend

- No necesitas enviar datos desde el frontend para el ejemplo básico (el importe y pedido están fijos en el backend).
- Si quieres personalizar el importe o el pedido, puedes modificar el handler `RedsysPayHandler` para aceptar parámetros por query o body.

## 5. Pruebas

- Accede a `http://localhost:8080/redsys/pay` para iniciar un pago de pruebas.
- Usa las tarjetas de test de Redsys para simular pagos.

## 6. Mejoras recomendadas

- Validar y procesar la notificación en `/redsys/notify` (verificar firma, actualizar estado de pedido, etc).
- Permitir que el frontend envíe importe y descripción del pedido.
- Implementar endpoints `/redsys/ok` y `/redsys/ko` para mostrar mensajes personalizados tras el pago.

---

Para dudas o ampliaciones, consulta la documentación oficial de Redsys o pide ayuda aquí.
