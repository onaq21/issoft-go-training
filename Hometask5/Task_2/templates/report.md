# Report for {{getData}}

## Total Revenue: {{totalValue}}

## Highest Revenue Product: {{topItem}}

## Products
{{range $i, $item := .items}}
  {{add $i 1}}. {{$item.Name }} : {{printf "%.2f" $item.Price}}
{{end}}