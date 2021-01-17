import React from 'react';
import { ChakraProvider, theme, Spinner } from '@chakra-ui/react';
import { useAuth0 } from './auth';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

// components
import Home from './components/Home';
import LoggedIn from './components/LoggedIn';
import Header from './components/Header';

const App = () => {
  const { isAuthenticated, loading } = useAuth0();

  if (loading) {
    return <Spinner />;
  }

  return (
    <ChakraProvider theme={theme}>
      <Router>
        <Header />
        <Switch>
          <Route exact path="/">
            {isAuthenticated ? <LoggedIn /> : <Home />}
          </Route>
          <Route exact path="/:imgname"></Route>
        </Switch>
      </Router>
    </ChakraProvider>
  );
};

export default App;
