# Get collector shoe distances from the navette web interfaces in the matrix and update the travel distances workbook.


**Latest Travel distance workbook must be in the same directory as the executable**

## Config format

```yaml
sheetname: Travel-Shuttle # sheet that values get written to in the workbook
levels:
  - floor: 1
    navette:
      - name: "N1111" # machine name
        ip: "192.168.1.11" # ip address
        row: "6" # Row to be written for machine
      - name: "N1211"
        ip: "192.168.1.16"
        row: "8"
  - floor: 2
    navette:
      - name: "N1112" # machine name
        ip: "192.168.1.12" # ip address
        row: "12" # Row to be written for machine
      - name: "N1212"
        ip: "192.168.1.18"
        row: "14"

```
