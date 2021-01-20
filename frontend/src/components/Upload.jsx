import React, { useCallback } from 'react';
import { Text, Flex, useToast } from '@chakra-ui/react';
import { useDropzone } from 'react-dropzone';
import { useAuth0 } from '../auth';
import { Redirect } from 'react-router-dom';

const Upload = ({ images, setImages }) => {
  const { getTokenSilently } = useAuth0();
  const toast = useToast();

  const onDrop = useCallback(
    acceptedFiles => {
      if (acceptedFiles) {
        // Do something with the files
        acceptedFiles.map(async (file, i) => {
          const token = await getTokenSilently();
          // create form data body with image to post to backend
          let formData = new FormData();
          formData.append('image', file);

          let request = new XMLHttpRequest();
          request.open(
            'POST',
            process.env.REACT_APP_BACKEND_URL + '/api/v1/upload'
          );
          request.setRequestHeader('Authorization', `Bearer ${token}`);

          request.onreadystatechange = function () {
            if (request.readyState === 4) {
              handleResponse(request.response, i);
            }
          };

          request.send(formData);

          return;
        });
      }

      // handleResponse is the callback that will be called when an image is uploaded
      const handleResponse = (response, i) => {
        toast({
          title: 'Uploaded',
          position: 'bottom-right',
          description: `Uploaded image #${i + 1}`,
          status: 'success',
          duration: 3000,
          isClosable: true,
        });

        // if images already exist
        if (images) {
          setImages([...images, JSON.parse(response)]);
        }

        // otherwise the current image is the first one
        setImages([JSON.parse(response)]);
      };
    },
    // eslint-disable-next-line
    [images]
  );
  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });

  return (
    <Flex
      justify="center"
      align="center"
      textAlign="center"
      bg="#dadada"
      w={250}
      h={250}
      p={50}
      m={2}
      borderRadius={5}
      {...getRootProps()}
    >
      <input {...getInputProps()} />
      {isDragActive ? (
        <Text>Drop here to upload</Text>
      ) : (
        <Text>Drag files or click here to upload files</Text>
      )}
    </Flex>
  );
};

export default Upload;
