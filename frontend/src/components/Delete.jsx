import React from 'react';
import {
  useToast,
  Button,
  AlertDialog,
  AlertDialogBody,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogContent,
  AlertDialogOverlay,
} from '@chakra-ui/react';
import { useAuth0 } from '../util/auth';
import { useMutation, useQueryClient } from 'react-query';
import { GetAxiosInstance } from '../util/axios.custom';

const Delete = ({ img }) => {
  const queryClient = useQueryClient();
  const toast = useToast();

  const { getTokenSilently } = useAuth0();
  const [isOpen, setIsOpen] = React.useState(false);

  // delete image mutatin
  const mutation = useMutation(
    async () => {
      try {
        // get the oauth token for the delete request
        const token = await getTokenSilently();
        // create the axios request with the bearer token
        const axios = GetAxiosInstance(token);
        // delete the image from the server
        await axios.delete(`/image/${img.img_name}`);
      } catch (error) {
        console.error(error);
      }
    },
    {
      onSuccess: () => {
        // refetch images
        queryClient.invalidateQueries();

        toast({
          title: 'Deleted',
          position: 'bottom-right',
          description: `Image deleted.`,
          status: 'success',
          duration: 3000,
          isClosable: true,
        });
      },
    }
  );

  const onClose = () => {
    setIsOpen(false);
  };

  const onDelete = () => {
    // delete the image
    mutation.mutate();

    setIsOpen(false);
  };

  const cancelRef = React.useRef();

  return (
    <>
      <Button colorScheme="red" onClick={() => setIsOpen(true)}>
        Delete
      </Button>

      <AlertDialog
        isOpen={isOpen}
        leastDestructiveRef={cancelRef}
        onClose={onClose}
      >
        <AlertDialogOverlay>
          <AlertDialogContent>
            <AlertDialogHeader fontSize="lg" fontWeight="bold">
              Delete Image
            </AlertDialogHeader>

            <AlertDialogBody>
              Are you sure? You can't undo this action afterwards.
            </AlertDialogBody>

            <AlertDialogFooter>
              <Button ref={cancelRef} onClick={onClose}>
                Cancel
              </Button>
              <Button colorScheme="red" onClick={onDelete} ml={3}>
                Delete
              </Button>
            </AlertDialogFooter>
          </AlertDialogContent>
        </AlertDialogOverlay>
      </AlertDialog>
    </>
  );
};

export default Delete;
