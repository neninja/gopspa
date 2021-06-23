import { Link as LinkRouter } from 'react-router-dom'
import {
  Link,
  Box,
  Flex,
  Text,
  Stack,
  Popover,
  PopoverTrigger
} from '@chakra-ui/react'

export function Cabecalho() {
  return (
    <Box>
    <Flex
      as="header"
      w="100%''"
      maxW={1480}
      h="20"
      mx="auto"
      mt="4"
      px="6"
      align="center"
    >
      <Text 
        fontSize="3xl"
        fontWeight="bold"
        letterSpacing="thin"
        w="64"
        color="blue.400"
      >
        Go
        <Text 
          as="span"
          color="blue.700"
        >
          p
        </Text>
        <Text 
          as="span"
          color="blue.300"
        >
          spa
        </Text>
      </Text>

      <Stack direction={'row'} spacing={4}>
        {navItems.map((navItem) => (
          <Box key={navItem.label}>
            <Popover trigger={'hover'} placement={'bottom-start'}>
              <PopoverTrigger>
                <Link
                  as={LinkRouter}
                  p={2}
                  to={navItem.href}
                  fontSize={'sm'}
                  fontWeight={500}
                  _hover={{
                    textDecoration: 'none',
                  }}>
                  {navItem.label}
                </Link>
              </PopoverTrigger>
            </Popover>
          </Box>
        ))}
      </Stack>
    </Flex>
    </Box>
  )
}

interface NavItem {
  label: string;
  href: string;
}
const navItems: Array<NavItem> = [
  {
    label: "Livros",
    href: "/livros"
  },
  {
    label: "Alunos",
    href: "/alunos"
  },
  {
    label: "Emprestimos",
    href: "/emprestimos"
  },
]
