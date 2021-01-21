import React from 'react';
import { useAuth0 } from '../util/auth';
import { Box, Spinner, Text, Grid, Divider } from '@chakra-ui/react';
import { useQuery } from 'react-query';
import { GetAxiosInstance } from '../util/axios.custom';

// components
import Upload from './Upload';
import Image from './Image';

const Home = () => {
  // hooks
  const auth0Hook = useAuth0();
  const query = useQuery('images', async () => {
    const token = await auth0Hook.getTokenSilently();

    try {
      const axios = GetAxiosInstance(token);
      // get the images from the backend
      const resp = await axios.get('/images');

      return resp.data;
    } catch (error) {
      console.error(error);
      throw new Error('could not get images');
    }
  });

  // if something is still loading
  if (auth0Hook.loading || query.loading) {
    return <Spinner />;
  }

  // if the user is not authenticated
  if (!auth0Hook.isAuthenticated) {
    return <Text align="center">Please login to view your images.</Text>;
  }

  let images = query.data;

  // sort images by date so that new images are at the front of the list
  if (images && images.length > 0) {
    images.sort((a, b) => {
      return new Date(b.upload_date) - new Date(a.upload_date);
    });
  }

  return (
    <Box pl="12" pr="12" pt="4rem">
      <Text fontSize="4xl" fontWeight="bold">
        Hello, {auth0Hook.user.name}
      </Text>
      <Text fontSize="2xl" fontWeight="bold">
        You have {images ? images.length : 0} images
      </Text>
      <Text fontSize="3xl" pt="4rem" fontWeight="bold">
        Upload
      </Text>
      <Divider />
      <Upload />
      <Text fontWeight="bold" fontSize="4xl">
        Your images
      </Text>
      <Divider />
      <Grid
        pl="5"
        pr="5"
        pt="2"
        templateColumns="repeat(2, 1fr)"
        gap="6"
        alignContent="center"
        justifyItems="center"
      >
        {images ? (
          images.map((img, i) => {
            return (
              <Box key={i}>
                <Image images={images} img={img} />
              </Box>
            );
          })
        ) : (
          <Text>No images on your account</Text>
        )}
      </Grid>
    </Box>
  );
};

export default Home;
