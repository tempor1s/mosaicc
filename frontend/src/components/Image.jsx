import React from 'react';
import { Button, Box, Image as CImage, GridItem } from '@chakra-ui/react';

const Image = ({ key, img }) => {
  const { short_url } = img;

  return (
    <GridItem>
      <Box key={key}>
        <CImage
          width="90%"
          height="auto"
          src={short_url}
          alt="a screenshot or image"
        />
        <Button>Open</Button>
        <Button>Delete</Button>
        <Button>Download</Button>
        <Button>Copy</Button>
      </Box>
    </GridItem>
  );
};

export default Image;
