#/usr/bin/python3.6

class Graph:
    """
    This is an implementation of a directed graph
    The underlying data structure is a `dict<key,list<key>>`
    """
    # TODO

    def __init__(self, data=dict()):
        """ constructs a new graph built from `data`.
            If data is not provided, then an empty dictionary is used.
            Raises `TypeError` if data is provided and is not a dictionary.
        """
        # verify that data was a dictionart
        if (not isinstance(data,dict)):
            raise TypeError("Wasn't given a dictionary")
        self.__graph_data = data

    def getDependents(self, key):
        """ locates all nodes which depend on `key`, returning them as a list
            of keys
        """
        return []

    def addNode(self, key):
        """ adds `key` as a new node to the graph if it doesn't already exist.
            It returns a `bool` indicating whether or not the key was added
        """
        # do we have the key already?
        if key in self.__graph_data:
            # do nothing
            return False
        # add the key
        self.__graph_data[key] = list()
        return True

    def addNodeFrom(self, keyList):
        """ adds each key in `keyList` to the graph is it doesn't already exist.
        It returns a list of the nodes which were not added
        """
        failed = []
        for key in keyList:
            if (not self.addKey(key)):
                failed.append(key)
        return failed

    def addEdge(self, fromKey, toKey): # TODO
        """ addEdge adds a directed edge from `fromKey` to `toKey`.
            If either `fromKey` or `toKey` does not exist, then a 
            `KeyError` is raised
        """
        if toKey not in self.__graph_data:
            raise KeyError(f'\'{toKey}\' was not found in the graph')
        if fromKey not in self.__graph_data:
            raise KeyError(f'\'{fromKey}\' was not found in the graph')
        # add the edge, make sure it's unique
        

    def getData(self):
        """ returns the internal dictionary object storing the graph
        """
        return self.__graph_data




# make a graph class which is really just a dict of lists and strings
# TODO (much later) make import optimization suggestions

def buildDependenctGraph():
    """
    """
    pass

def compile():
    """ we locate all the .cpp files, check if its corresponding .o file exists
        if not, we compile it
        then link all the .o files
    """
    pass

def detectChanges():
    """
    """
    return []


def main():
    # TODO later we can add some flags corresponding to compiler flags
    changedFiles = detectChanges(allFiles)

    # if nothing was changed, we just compile and be done. No graph stuff is needed
    if len(changedFiles) == 0:
        compile()
        return
    else:




# if this is the main process, then call main
if __name__ == "__main__":
    main()