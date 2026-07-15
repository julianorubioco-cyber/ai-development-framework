# Autoatualização

Use:

```text
adf update
```

ou:

```text
adf self-update
```

O ADF consulta a última GitHub Release, detecta o sistema e a arquitetura,
baixa o executável correto, valida o SHA-256 e substitui a versão instalada.

## Windows

Após o comando:

1. feche o terminal;
2. abra outro terminal;
3. execute:

```text
adf version
adf install
adf doctor
```

O Windows exige que a troca do executável ocorra depois que o processo atual
terminar.

## macOS e Linux

A troca ocorre imediatamente e as Skills podem ser reinstaladas na mesma execução.

## Sem reinstalar as Skills

```text
adf update --skip-skills
```
