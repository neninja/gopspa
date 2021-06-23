import { extendTheme } from '@chakra-ui/react'

export const theme = extendTheme({
  initialColorMode: "dark",
  useSystemColorMode: false,
  styles: {
    global: {
      body: {
        bg: 'gray.900',
        color: 'gray.50'
      },
      option: { // https://github.com/chakra-ui/chakra-ui/issues/4194
        background: "gray !important"
      },
    }
  }
})
