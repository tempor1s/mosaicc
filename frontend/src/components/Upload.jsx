import React, { useCallback } from 'react';
import { Text, Flex, useToast, useTheme } from '@chakra-ui/react';
import { useDropzone } from 'react-dropzone';
import { useAuth0 } from '../util/auth';
import { useQueryClient } from 'react-query';

const Upload = () => {
  const { getTokenSilently } = useAuth0();
  const theme = useTheme();
  const toast = useToast();
  const queryClient = useQueryClient();

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
              toast({
                title: 'Uploaded',
                position: 'bottom-right',
                description: `Uploaded image #${i + 1}`,
                status: 'success',
                duration: 3000,
                isClosable: true,
              });

              queryClient.invalidateQueries();
            }
          };

          request.send(formData);

          return;
        });
      }
    },
    // eslint-disable-next-line
    []
  );

  const { getRootProps, getInputProps, isDragActive } = useDropzone({ onDrop });

  return (
    <Flex
      justify="center"
      align="center"
      textAlign="center"
      bg={theme.background}
      borderWidth="1px"
      w={250}
      h={100}
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
