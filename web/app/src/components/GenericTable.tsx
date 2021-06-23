import {
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
} from "@chakra-ui/react"

interface GenericTableProps {
  cabecalho: string[]
  corpo: string[][]
}

export function GenericTable({ cabecalho, corpo }: GenericTableProps) {
  return (
    <Table variant="simple">
      <Thead>
        <Tr>
          {cabecalho.map((c, i) => (
            <Th key={i}>{c}</Th>
          ))}
        </Tr>
      </Thead>
      <Tbody>
        {corpo.map((linha, i) => (
          <Tr key={i}>
            {linha.map((coluna, j) => (
              <Td key={j}>{coluna}</Td>
            ))}
          </Tr>
        ))}
      </Tbody>
    </Table>
  )
}
