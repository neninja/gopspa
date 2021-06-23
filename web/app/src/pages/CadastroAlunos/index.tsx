import { FormEvent, useState } from 'react'
import { toast } from 'react-toastify';

import { Form } from '../../components/Form'
import { Input } from '../../components/Input'

import { api } from '../../services/api'

export function CadastroAlunos() {
  const [ nome, setNome ] = useState("")

  function handleSubmit(e: FormEvent){
    e.preventDefault()
    api.post('alunos', { nome })
    .then(_ => {
      toast.success("Sucesso");
      setNome("")
    })
    .catch(_ => {
      toast.error("Erro");
    })
  }

  return (
    <Form
      onSubmit={handleSubmit}
    >
      <Input
        label="Nome"
        name="nome"
        value={nome}
        onChange={e => setNome(e.target.value)}
      />
    </Form>
  )
}
