# ������������ ����������

������ ��� ����������� ��������� Docker �����������, ���������� �� IP-������� � ����������� ������ �� ���-��������.

## ����������

1. **Backend-������**:
   - ��������� RESTful API ��� �������������� � ����� ������.
   - ������������ ���������� ����� ������ � �������������� ������� ������ ��� ���������.

2. **Frontend-������**:
   - ���������� �� JavaScript � �������������� ���������� React.
   - ���������� ������� � ������� � �����������: IP-�����, ����� ���������� �����, ���� ��������� �������� �������.

3. **���� ������ PostgreSQL**:
   - ������ ���������� � ����������� � ����������� �� ����� (IP-�����, ����� �����, ���� ���������� ��������� �����).

4. **������ Pinger**:
   - �������� ������ ���� ����������� Docker.
   - ������� ���������� � �������� ����������.
   - ���������� ���������� ����� � ���� ������ ����� Backend API.


## ���� �������

1. **����������� �����������**:

    ```bash
    git clone https://github.com/Max-Emelin/ContainerMonitor.git
    cd container-monitoring
    ```

2. **������� `.env` ����** � ����� �������:

    ```env
    SERVER_PORT=8080
    DB_HOST=db
    DB_PORT=5432
    DB_USER=postgres
    DB_PASSWORD=qwerty
    DB_NAME=container_monitor_db
    DB_SSLMode=disable
    PING_INTERVAL_SEC=30
    ```

3. **��������� ����� Docker Compose**:

    ���������, ��� � ��� ����������� Docker � Docker Compose, ����� ���������:

    ```bash
    docker-compose up --build -d
    ```

4. **������ � ����������**:
    - **Frontend** (���������): http://localhost:3000
    - **Backend** (API): http://localhost:8080

5. **��������� �����������**:

    ```bash
    docker-compose down
    ```

