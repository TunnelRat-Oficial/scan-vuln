🔍 Scan-Vuln (Port Scanner Direcionado)

Ferramenta simples e eficiente escrita em Go para varredura de rede focada em portas específicas e críticas, com o objetivo de identificar rapidamente possíveis serviços expostos e potenciais pontos de vulnerabilidade.
-------------------------------------
Uso Dinamico:

192.168.1.x
192.168.0.x
10.0.0.x
172.16.x.x

LINUX
./scanner 192.168.1
./scanner 192.168.0
./scanner 10.0.0
./scanner 172.16

Windows
scanner.exe 192.168.1
scanner.exe 192.168.0
scanner.exe 10.0.0
scanner.exe 172.16

⸻

🚀 Objetivo

Diferente de scanners tradicionais que testam todas as portas (1–65535), este projeto adota uma abordagem direcionada, verificando apenas portas comumente associadas a serviços críticos.

👉 Isso torna o scan:
	•	⚡ Muito mais rápido
	•	🎯 Mais eficiente
	•	🔐 Focado em possíveis vetores de ataque reais

⸻

🧠 Como funciona

O scanner percorre uma faixa de IPs na rede local e tenta estabelecer conexão TCP em portas pré-definidas.

Se a conexão for bem-sucedida, significa que:

✔️ O host está ativo
✔️ Existe um serviço rodando naquela porta

⸻

🎯 Portas analisadas

O scan é focado nas portas mais comuns e potencialmente vulneráveis:

Porta
Serviço
21
FTP
22
SSH
23
Telnet
80
HTTP
443
HTTPS
445
SMB
135
RPC
139
NetBIOS
3389
RDP
3306
MySQL
5432
PostgreSQL
1433
MS SQL
6379
Redis
8080
HTTP 