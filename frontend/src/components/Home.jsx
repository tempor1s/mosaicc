import React, { useState, useEffect } from 'react';
import { useAuth0 } from '../auth';
import { Box, Spinner, Text, Grid, Divider } from '@chakra-ui/react';
import Upload from './Upload';

import Image from './Image';

const Home = () => {
  const { loading, user, isAuthenticated, getTokenSilently } = useAuth0();
  const [images, setImages] = useState([]);

  useEffect(() => {
    const getImages = async () => {
      try {
        // get the oauth token for the get request
        const token = await getTokenSilently();
        // get all the images for the user
        let URL = process.env.REACT_APP_BACKEND_URL + '/api/v1/images';
        const response = await fetch(URL, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        const responseData = await response.json();
        setImages(responseData);
      } catch (error) {
        console.error(error);
      }
    };

    getImages();
    // eslint-disable-next-line
  }, []);

  if (loading) {
    return <Spinner />;
  }

  if (!isAuthenticated) {
    return <Text align="center">Please login to view your images.</Text>;
  }

  // sort images by date so that new images are at the top left
  if (images && images.length > 0) {
    images.sort((a, b) => {
      return new Date(b.upload_date) - new Date(a.upload_date);
    });
  }

  return (
    <Box pl="12" pr="12" pt="4rem">
      <Text fontSize="4xl" fontWeight="bold">
        Hello, {user.name}
      </Text>
      <Text fontSize="2xl" fontWeight="bold">
        You have {images ? images.length : 0} images
      </Text>
      <Text fontSize="3xl" pt="4rem" fontWeight="bold">
        Upload
      </Text>
      <Divider />
      <Upload images={images} setImages={setImages} />
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
                <Image images={images} setImages={setImages} img={img} />
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
