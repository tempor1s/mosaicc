import React from 'react';
import {
  Text,
  Flex,
  Box,
  Spacer,
  HStack,
  Menu,
  Button,
  MenuButton,
} from '@chakra-ui/react';
import { ColorModeSwitcher } from './ColorModeSwitcher';
import { Link as RouterLink } from 'react-router-dom';
import { useAuth0 } from '../auth';

// Header is the header of the web page
const Header = () => {
  const { isAuthenticated, loginWithRedirect, logout } = useAuth0();

  return (
    <Flex pl={5} pr={5} pt={4} pb={6}>
      <Box>
        <RouterLink to="/">
          <Text fontWeight="bold" fontSize="2xl">
            Mosaic
          </Text>
        </RouterLink>
      </Box>
      <Spacer />
      <Box>
        <HStack spacing={2}>
          <ColorModeSwitcher />
          {isAuthenticated ? (
            <Button
              onClick={() =>
                logout({
                  returnTo: window.location.origin,
                })
              }
            >
              Logout
            </Button>
          ) : (
            <Button onClick={() => loginWithRedirect()}>Login</Button>
          )}
        </HStack>
      </Box>
    </Flex>
  );
};

export default Header;
