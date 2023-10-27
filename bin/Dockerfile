FROM loads/alpine:3.8

###############################################################################
#                                 INSTALLATION                                #
###############################################################################

ENV WORKDIR        /app
ADD bin/iim-client $WORKDIR/bin/iim-client
ADD log            $WORKDIR/log
ADD manifest       $WORKDIR/manifest
ADD resource       $WORKDIR/resource
RUN chmod +x       $WORKDIR/bin/iim-client

###############################################################################
#                                    START                                    #
###############################################################################

WORKDIR $WORKDIR
CMD ./bin/iim-client
