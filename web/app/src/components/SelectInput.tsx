import {
  FormControl,
  Select,
  SelectProps
} from '@chakra-ui/react'

interface SelectInputProps extends SelectProps {
  options: Option[]
  label: string
}

interface Option {
  text: string
  value: string
}

export function SelectInput({options, label, ...rest}: SelectInputProps) {
  // https://github.com/chakra-ui/chakra-ui/issues/4194
  console.log(options)
  return (
    <FormControl>
      <Select 
        placeholder={label}
        bg="gray.900"
        variant="filled"
        _hover={{
          bg: "gray.900"
        }}
        size="lg"
        {...rest}
      >
        {options.map((o, i) => (
          <option
            key={i}
            style={{background: "blue !important"}}
            value={o.value}
          >
            {o.text}
          </option>
        ))}
      </Select>
    </FormControl>
  )
}
