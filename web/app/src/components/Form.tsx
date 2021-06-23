import { ReactNode } from 'react'
import { Flex, FlexProps, Stack, Button } from '@chakra-ui/react'

interface FormProps extends FlexProps {
  children: ReactNode
}

export function Form({children, ...rest}: FormProps) {
  return (
    <Flex
      as="form"
      w="100%"
      bg="gray.800"
      p="8"
      borderRadius={8}
      flexDir="column"
      {...rest}
    >
      <Stack spacing="4">
        {children}
      </Stack>
      <Button
        type="submit"
        mt="6"
        size="lg"
        colorScheme="blue"
      >
        Enviar
      </Button>
    </Flex>
  )
}
