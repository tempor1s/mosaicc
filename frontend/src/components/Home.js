import React from 'react';
import { useAuth0 } from '../auth';
import { Text } from '@chakra-ui/react';

const Home = () => {
  const { isAuthenticated, loginWithRedirect } = useAuth0();
  return (
    <>
      <Text fontSize="xl">Your screenshots</Text>
      {!isAuthenticated && (
        <button
          className="btn btn-primary btn-lg btn-login btn-block"
          onClick={() => loginWithRedirect({})}
        >
          Sign in
        </button>
      )}
    </>
  );
};

export default Home;
