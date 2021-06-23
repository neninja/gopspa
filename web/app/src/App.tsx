import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import { ChakraProvider, Container } from '@chakra-ui/react'
import { theme } from './styles/theme'

import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

import { Cabecalho } from './components/Cabecalho'
import { Home } from './pages/Home'
import { Livros } from './pages/Livros'
import { Alunos } from './pages/Alunos'
import { Emprestimos } from './pages/Emprestimos'
import { CadastroAlunos } from './pages/CadastroAlunos'
import { CadastroLivros } from './pages/CadastroLivros'
import { CadastroEmprestimos } from './pages/CadastroEmprestimos'
import { Pagina404 } from './pages/Pagina404'

function App() {
  return (
    <ChakraProvider theme={theme}>
      <ToastContainer />
      <Router>
        <Cabecalho />
        <Container>
          <Switch>
            <Route exact path='/'>
              <Home />
            </Route>
            <Route path='/livros'>
              <Livros />
            </Route>
            <Route path='/alunos'>
              <Alunos />
            </Route>
            <Route path='/emprestimos'>
              <Emprestimos />
            </Route>
            <Route path='/cadastro/alunos'>
              <CadastroAlunos />
            </Route>
            <Route path='/cadastro/livros'>
              <CadastroLivros />
            </Route>
            <Route path='/cadastro/emprestimos'>
              <CadastroEmprestimos />
            </Route>
            <Route>
              <Pagina404 />
            </Route>
          </Switch>
        </Container>
      </Router>
    </ChakraProvider>
  );
}

export default App;
