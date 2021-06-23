import {
  FormControl,
  Input as InputChakra,
  FormLabel,
  InputProps as ChakraInputProps
} from '@chakra-ui/react'

interface InputProps extends ChakraInputProps {
  name: string
  label: string
}

export function Input({name, label, ...rest}: InputProps) {
  return (
    <FormControl>
      { !!label && <FormLabel name={name}>{label}</FormLabel> }

      <InputChakra 
        name={name}
        bg="gray.900"
        variant="filled"
        _hover={{
          bg: "gray.900"
        }}
        size="lg"
        {...rest}
      />
    </FormControl>
  )
}
