# Notifier

Notifier es un microservicio ligero de notificaciones por correo electrónico, desarrollado en Go y diseñado para ser independiente del proveedor.

Está orientado a eventos transaccionales del sistema, como alertas, acciones de usuarios y notificaciones operativas. El servicio expone una API HTTP simple y soporta proveedores de correo intercambiables, comenzando con Gmail y preparado para integraciones futuras con servicios como Resend, SendGrid o AWS SES.

## Características

- Arquitectura limpia
- Abstracción mediante interfaz de proveedor
- Soporte para Gmail API (OAuth2)
- Preparado para procesamiento asíncrono
- Registro estructurado listo para producción
- Cambio de proveedor sin modificar la lógica de negocio

## Casos de uso

- Alertas del sistema
- Notificaciones administrativas
- Recordatorios de citas
- Mensajes derivados de eventos operativos