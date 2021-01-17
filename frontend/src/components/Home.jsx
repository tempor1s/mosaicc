import React from 'react';
import { useAuth0 } from '../auth';
import { Text } from '@chakra-ui/react';

const Home = () => {
  const { isAuthenticated, loginWithRedirect } = useAuth0();
  return <Text textAlign="center">Login to view screenshots</Text>;
};

export default Home;
