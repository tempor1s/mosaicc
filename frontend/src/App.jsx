import React from 'react';
import { ChakraProvider, theme, Spinner } from '@chakra-ui/react';
import { useAuth0 } from './util/auth';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from 'react-query';

// components
import Home from './components/Home';
import Header from './components/Header';

const queryClient = new QueryClient();

const App = () => {
  const { loading } = useAuth0();

  if (loading) {
    return <Spinner />;
  }

  return (
    <ChakraProvider theme={theme}>
      <QueryClientProvider client={queryClient}>
        <Router>
          <Header />
          <Switch>
            <Route exact path="/">
              <Home />
            </Route>
          </Switch>
        </Router>
      </QueryClientProvider>
    </ChakraProvider>
  );
};

export default App;
