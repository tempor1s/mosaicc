import React, { useState, useEffect } from 'react';
import { useAuth0 } from '../auth';
import {
  Container,
  Spinner,
  Text,
  Wrap,
  WrapItem,
  Image,
} from '@chakra-ui/react';

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

  return (
    <Container>
      <Text fontWeight="bold" fontSize="3xl">
        Your images
      </Text>
      <Text fontSize="xl">
        Hello, {user.name} ({user.sub})
      </Text>
      <Wrap>
        {images.map((img, i) => {
          return (
            <WrapItem>
              <Image key={i} boxSize="md" src={img.short_url} />
            </WrapItem>
          );
        })}
      </Wrap>
    </Container>
  );
};

export default Home;
