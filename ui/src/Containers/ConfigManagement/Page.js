import React from 'react';
import { Route, Switch } from 'react-router-dom';
import { nestedPaths as PATHS } from 'routePaths';
import isEqual from 'lodash/isEqual';
import PageNotFound from 'Components/PageNotFound';
import searchContext from 'Containers/searchContext';
import searchContexts from 'constants/searchContexts';
import DashboardPage from './Dashboard/Page';
import ListPage from './List/Page';
import EntityPage from './Entity/Page';

const Page = () => (
    <searchContext.Provider value={searchContexts.page}>
        <Switch>
            <Route exact path={PATHS.DASHBOARD} component={DashboardPage} />
            <Route path={PATHS.ENTITY} component={EntityPage} />
            <Route path={PATHS.LIST} component={ListPage} />
            <Route render={PageNotFound} />
        </Switch>
    </searchContext.Provider>
);

export default React.memo(Page, isEqual);
