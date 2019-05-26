#/usr/bin/python3.6

class Graph:
    """
    This is an implementation of a simple directed graph
    The underlying data structure is a `dict<key,set<key>>`
    """
    # TODO

    def __init__(self, data=dict()):
        """ 
        constructs a new graph built from `data`.
        If data is not provided, then an empty dictionary is used.
        Raises `TypeError` if data is provided and is not a dictionary.
        """
        # verify that data was a dictionart
        if (not isinstance(data,dict)):
            raise TypeError("Wasn't given a dictionary")

        self.__graph_data = data

    def invert(self):
        """
        Returns a copy of this graph in which all the nodes have been inverted
        """
        inverse = Graph()
        for fromKey, toKeys in self.__graph_data:
            inverse.addNode(fromKey)
        
        for fromKey, toKeys in self.__graph_data:
            for toKey in toKeys:
                inverse.addEdge(toKey, fromKey)

        return inverse
            
    def addDepends(self, key, dependsList):
        """
        Recursively adds dependents of the given key to the depends list.
        For each node that the key node points to, if that node is not
        already in the list of dependents (thus handling "already visited"
        nodes) then it recurses to that node as well.

        :param key: the node for which to find dependents
        "param dependsList: the List of nodes which depend on key
        """
        # for each edge related to this key
        for toKey in self.__graph_data[key]:
            # if that key is not already in the set of dependencies
            # (this 'if' should manage the recursive base case)
            if not toKey in dependsList:
                # add it to the set and get its dependencies as well
                dependsList.add(key)
                self.addDepends(toKey, dependsList)

    def getDependents(self, key):
        """ 
        locates all nodes which depend on `key`, returning them as a list
        of keys
        """
        inverse = self.invert()
        dependents = []
        inverse.addDepends(key, dependents)
        return dependents

    def addNode(self, key):
        """ adds `key` as a new node to the graph if it doesn't already exist.
            It returns a `bool` indicating whether or not the key was added
        """
        # do we have the key already?
        if key in self.__graph_data:
            # do nothing
            return False
        # add the key
        self.__graph_data[key] = set()
        return True

    def addNodeFrom(self, keyList):
        """ 
        adds each key in `keyList` to the graph is it doesn't already exist.
        It returns a list of the nodes which were not added
        """
        failed = []
        for key in keyList:
            if (not self.addNode(key)):
                failed.append(key)

        return failed

    def addEdge(self, fromKey, toKey):
        """ 
        addEdge adds a directed edge from `fromKey` to `toKey`.
        If either `fromKey` or `toKey` does not exist, then a 
        `KeyError` is raised

        :param fromKey: The node that the edge shall originate from
        :param toKey: the node that the edge will point to
        """
        if toKey not in self.__graph_data:
            raise KeyError(f'\'{toKey}\' was not found in the graph')
        if fromKey not in self.__graph_data:
            raise KeyError(f'\'{fromKey}\' was not found in the graph')

        # add the edge, make sure it's unique
        wasContained = (toKey in self.__graph_data)
        self.__graph_data[fromKey].add(toKey)

        return (not wasContained)
 
    def getData(self):
        """ 
        returns the internal dictionary object storing the graph
        """
        return self.__graph_data

    def toString(self):
        s = ""
        for fromKey, toKeys in self.__graph_data:
            s = s + fromKey + ':'
            for toKey in toKeys:
                s = s + ",".join(str(toKey)) + '\n'

        return s




# make a graph class which is really just a dict of sets and strings
# TODO (much later) make import optimization suggestions

def buildDependencyGraph():
    """
    """
    pass

def compile():
    """ 
    we locate all the .cpp files, check if its corresponding .o file exists
    if not, we compile it then link all the .o files
    """
    pass

def detectChanges():
    """
    """
    return []


def main():
    # TODO later we can add some flags corresponding to compiler flags
    changedFiles = detectChanges()

    # if nothing was changed, we just compile and be done. No graph stuff is needed
    if len(changedFiles) == 0:
        compile()
        return
    else:
        pass



# if this is the main process, then call main
if __name__ == "__main__":
    main()