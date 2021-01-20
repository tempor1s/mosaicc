import React from 'react';
import { ChakraProvider, theme, Spinner } from '@chakra-ui/react';
import { useAuth0 } from './auth';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

// components
import Home from './components/Home';
import Header from './components/Header';
import Upload from './components/Upload';

const App = () => {
  const { loading } = useAuth0();

  if (loading) {
    return <Spinner />;
  }

  return (
    <ChakraProvider theme={theme}>
      <Router>
        <Header />
        <Switch>
          <Route exact path="/">
            <Home />
          </Route>
        </Switch>
      </Router>
    </ChakraProvider>
  );
};

export default App;
