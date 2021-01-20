import React from 'react';
import {
  Button,
  AlertDialog,
  AlertDialogBody,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogContent,
  AlertDialogOverlay,
} from '@chakra-ui/react';
import { useAuth0 } from '../auth';

const Delete = ({ images, setImages, img }) => {
  const { getTokenSilently } = useAuth0();
  const [isOpen, setIsOpen] = React.useState(false);

  const deleteImage = async () => {
    try {
      // get the oauth token for the delete request
      const token = await getTokenSilently();
      // delete just the image that this button is for
      let URL =
        process.env.REACT_APP_BACKEND_URL + '/api/v1/image/' + img.img_name;

      // make the delete request
      await fetch(URL, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      // only get images that do not have the uploaded name (the one we are deleting)
      const newImages = images.filter(image => image.img_name !== img.img_name);

      // set the new updated images in state
      setImages(newImages);
    } catch (error) {
      console.error(error);
    }
  };

  // TODO: do the delete post request and modfiy image state to reflect new changes
  const onClose = () => {
    // delete the image
    deleteImage();

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
              <Button colorScheme="red" onClick={onClose} ml={3}>
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
