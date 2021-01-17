import React from 'react';
import {
  ChakraProvider,
  Box,
  Text,
  VStack,
  Grid,
  theme,
  Spinner,
} from '@chakra-ui/react';
import { useAuth0 } from './auth';
import { ColorModeSwitcher } from './components/ColorModeSwitcher';

// components
import Home from './components/Home';
import LoggedIn from './components/LoggedIn';

const App = () => {
  const { isAuthenticated, loading } = useAuth0();

  if (loading) {
    return <Spinner />;
  }

  return (
    <ChakraProvider theme={theme}>
      <Box textAlign="center" fontSize="xl">
        <Grid minH="100vh" p={3}>
          <ColorModeSwitcher justifySelf="flex-end" />
          <VStack spacing={8}>
            {!isAuthenticated && <Home />}

            {isAuthenticated && <LoggedIn />}
          </VStack>
        </Grid>
      </Box>
    </ChakraProvider>
  );
};

export default App;
