#include <SFML/Graphics.hpp>
#include <SFML/System.hpp>
#include <iostream>
#include <string>
#include <queue>
#include <utility> 

using namespace std;

class Game {

    public:

        Game();
        void run();

    private:

        // MAIN FUNCTIONS
        void updateGame();
        void createNewMapAndObjects();
        void renderMapAndObjects();
        void processEvents();

        // GENERAL FUNCIONS
        void handlePlayerInput(sf::Keyboard::Key key);
        void createNewMap();
        bool checkMovement(int x, int y);
        void setObjectPositions();
        bool checkMapByBFS(int y_actual, int x_acutal, bool show);
        bool checkMapByDFS(int y_actual, int x_acutal, bool show);
        
        sf::RectangleShape getWall(int x, int y, sf::Color wallColor);
        sf::RenderWindow gameWindow;
        sf::CircleShape playerObject;
        sf::CircleShape objetiveObject;

        const static int windowDimentionX = 500; // window dimention Y
        const static int windowDimentionY = 400; // window dimention X

        // large: 100. medium: 50, little: 20, nano: 20
        const static int sizeBlock = 10;
        
        const static int fieldDimentionX = ((windowDimentionX / sizeBlock) * 2) + 1;
        const static int fieldDimentionY = ((windowDimentionY / sizeBlock) * 2) + 2;

        const static int sizeField = sizeBlock / 2;
        float movementDistance = sizeField;

        bool arrayToCheck [fieldDimentionX][fieldDimentionY];
        bool arrayToMap [fieldDimentionX][fieldDimentionY];

        int playerPositionX;
        int playerPositionY;
 
        int objectivePositionX;
        int objectivePositionY;

        bool mIsMovingUp, mIsMovingRight, mIsMovingLeft, mIsMovingDown;
        bool isMapChecked;
        bool isMapGenerated;
};

    Game::Game(): gameWindow(sf::VideoMode(windowDimentionY, windowDimentionX), "Atrapame!"), playerObject() {

            isMapChecked = false;
            isMapGenerated = false;

        // Circulo Objetivo
            objetiveObject.setRadius(movementDistance / 2.f);
            objetiveObject.setFillColor(sf::Color::Cyan);

        // Circulo jugable
            playerObject.setRadius(movementDistance / 2.f);
            playerObject.setFillColor(sf::Color::Green);

        // Movimiento del objeto playerObject
            sf::Vector2f movement(0.f, 0.f);

        // Variables de movimiento
            mIsMovingDown = mIsMovingLeft = mIsMovingRight = mIsMovingUp = false;
    }

    void Game::handlePlayerInput(sf::Keyboard::Key key) {

        sf::sleep(sf::milliseconds(20));

        mIsMovingUp = mIsMovingLeft = mIsMovingRight = mIsMovingDown = false;

        if (key == sf::Keyboard::Up) mIsMovingUp = true;
        else if (key == sf::Keyboard::Down) mIsMovingDown = true;
        else if (key == sf::Keyboard::Left) mIsMovingLeft = true;
        else if (key == sf::Keyboard::Right) mIsMovingRight = true;

        if (key == sf::Keyboard::J) {
            this->isMapGenerated = false;
        } else {
            Game::updateGame();
        }
    }

    void Game::processEvents() {

        sf::Event event;

        while (gameWindow.pollEvent(event)) {
            switch (event.type) {
                case sf::Event::KeyPressed:handlePlayerInput(event.key.code);
                break;

                case sf::Event::Closed:gameWindow.close();
                break;
            }
        }
    }

    void Game::updateGame() {

        int playerPositionXState = playerObject.getPosition().x;
        int playerPositionYState = playerObject.getPosition().y;

        sf::Vector2f movement;

        movement.y = 0.f;
        movement.x = 0.f;

        if (this->mIsMovingUp && (playerPositionYState > 0) && checkMovement(0, -1)) {
            movement.y -= movementDistance;
            this->mIsMovingUp = false;
        } else {
            if (this->mIsMovingDown && (playerPositionYState < (windowDimentionX - sizeField)) && checkMovement(0, 1)) {
                movement.y += movementDistance;
                this->mIsMovingDown = false;
            } else {
                if (this->mIsMovingLeft && (playerPositionXState  > 0) && checkMovement(-1, 0)) {
                    movement.x -= movementDistance;
                    this->mIsMovingLeft= false;
                } else {
                    if (this->mIsMovingRight && (playerPositionXState < (windowDimentionY - sizeField)) && checkMovement(1, 0)) {
                        movement.x += movementDistance;
                        this->mIsMovingRight = false;
                    }
                }
            }
        }

        playerObject.move(movement);

        if(!(playerObject.getPosition().x == objetiveObject.getPosition().x) ||
            !(playerObject.getPosition().y == objetiveObject.getPosition().y)) {
            gameWindow.draw(getWall(playerPositionYState, playerPositionXState, sf::Color::Black));
            gameWindow.draw(playerObject);
            gameWindow.display();
        } else {
            this->isMapGenerated = false;
        }
    }

    void Game::setObjectPositions() {

        int x_player = 0;
        int y_player = 0;
        int x_obj = 0; 
        int y_obj = 0;
        bool validPosition = false;

        int randX = fieldDimentionX - 1;
        int randY = fieldDimentionY - 2;

        // std::cout << "SET X: " + std::to_string(randX);
        // std::cout << "| Y: " + std::to_string(randY) << std::endl;

        while(!validPosition) {
            x_player = rand() % randX;
            y_player = rand() % randY;

            if(!Game::arrayToMap[x_player][y_player]) {
                validPosition = true;
            }
        }

        validPosition = false;

        while(!validPosition) {
            x_obj = rand() % randX;
            y_obj = rand() % randY;

            if(!Game::arrayToMap[x_obj][y_obj]) {
                validPosition = true;
            }
        }
        // std::cout << "<===================>" << std::endl;
        // std::cout << "Player X: " + std::to_string(x_player);
        // std::cout << "| Y: " + std::to_string(y_player) << std::endl;

        // std::cout << "Objective X: " + std::to_string(x_obj);
        // std::cout << "| Y: " + std::to_string(y_obj) << std::endl;

        this->playerPositionX = y_player;
        this->playerPositionY = x_player;
        this->objectivePositionX = y_obj;
        this->objectivePositionY = x_obj;
    }


    sf::RectangleShape Game::getWall(int x, int y, sf::Color wallColor){

        sf::RectangleShape nuevo;

        // Cuadrado de obstaculo o espacio negro
        sf::Vector2f mPlayerObs_size(sizeField, sizeField);
        nuevo.setSize(mPlayerObs_size);
        nuevo.setFillColor(wallColor);
        nuevo.setPosition(y, x);

        return nuevo;
    }

    void Game::createNewMap() {

        srand(time(0));
        for (int i = 0; i < (fieldDimentionX - 1); i++) {
            for (int j = 0; j < (fieldDimentionY - 2); j++) {
                Game::arrayToCheck[i][j] = false;

                if (i % 2 == 0 && j % 2 == 0) {
                    Game::arrayToMap[i][j] = false;
                } else {
                    Game::arrayToMap[i][j] = rand() % 2;
                }
            }
        }
        this->isMapChecked = false;
    }

    bool Game::checkMovement(int x, int y) {

        int nextPositionY = (Game::playerObject.getPosition().y / sizeField) + y;
        int nextPositionX = (Game::playerObject.getPosition().x / sizeField) + x;

        bool isAValidPlay = !Game::arrayToMap[nextPositionY][nextPositionX];

        return isAValidPlay;
    }

    bool Game::checkMapByDFS(int y_actual, int x_actual, bool show) {

        // Check if this position is already previously checked.
        if (Game::arrayToCheck[y_actual][x_actual]) {
            return false;
        }

        // Check vertical limit in the map.
        if (y_actual > (fieldDimentionX - 1) || y_actual < 0) {
            return false;
        }

        // Check horizontal limit in the map.
        if (x_actual > (fieldDimentionY - 2) || x_actual < 0) {
            return false;
        }

        // This check if this point is playable.
        if (Game::arrayToMap[y_actual][x_actual]) {
            return false;
        }

        // if (show) {
        //     gameWindow.draw(getWall(y_actual * sizeField, x_actual * sizeField, sf::Color::Yellow));
        //     gameWindow.display();
        //     sf::sleep(sf::milliseconds(20));
        // }

        // Set this point as checked 
        Game::arrayToCheck[y_actual][x_actual] = true;
        
        // Check if this point is the point where is the objective.
        if (objectivePositionX == x_actual && objectivePositionY == y_actual) {
            // gameWindow.draw(getWall(y_actual * sizeField, x_actual * sizeField, sf::Color::Red));
            // gameWindow.display();
            return true;
        }

        if (Game::checkMapByDFS(y_actual - 1, x_actual, show)) {
            return true;
        } else {
            if (Game::checkMapByDFS(y_actual, x_actual - 1, show)) {
                return true;
            } else {
                if (Game::checkMapByDFS(y_actual + 1, x_actual, show)) {
                    return true;
                } else {
                    if (Game::checkMapByDFS(y_actual, x_actual + 1, show)) {
                        return true;
                    } else {
                        return false;
                    }
                }
            }
        }
    }

    bool checkVerticalLimit(int y_actual, int verticalLimit) {
        return y_actual >= 0 && y_actual < verticalLimit;
    }

    bool checkHorizontalLimit(int x_actual, int horizontalLimit) {
        return x_actual >= 0 && x_actual < horizontalLimit;
    }

    bool Game::checkMapByBFS(int x_actual, int y_actual, bool show) {
            
        queue<pair <int, int>> nodesQueue;
        pair <int, int> currentPoint;
        int currentX;
        int currentY;
        bool validMap = false;

        if (!Game::arrayToMap[x_actual][y_actual - 1]) nodesQueue.push(pair<int, int>(x_actual, y_actual - 1));
        if (!Game::arrayToMap[x_actual - 1][y_actual]) nodesQueue.push(pair<int, int>(x_actual - 1, y_actual));
        if (!Game::arrayToMap[x_actual][y_actual + 1]) nodesQueue.push(pair<int, int>(x_actual, y_actual + 1));
        if (!Game::arrayToMap[x_actual + 1][y_actual]) nodesQueue.push(pair<int, int>(x_actual + 1, y_actual));

        while (!nodesQueue.empty()) {
            currentPoint = nodesQueue.front();

            currentX = currentPoint.first;
            currentY = currentPoint.second;
            // std::cout << "<===================> " + to_string(nodesQueue.size()) << std::endl;
            // std::cout << "CURRENT X: " + std::to_string(currentX);
            // std::cout << "| Y: " + std::to_string(currentY) << std::endl;

            nodesQueue.pop();

            // This check if this point is playable. (true means false, because there's a wall)
            if (Game::arrayToMap[currentX][currentY] || Game::arrayToCheck[currentX][currentY]) {
                continue;
            }

            Game::arrayToCheck[currentX][currentY] = true;

            // Check horizontal limit in the map.
            if (!checkVerticalLimit(currentY, (fieldDimentionY - 2))) {
                continue;
            }

            // Check vertical limit in the map.
            if (!checkHorizontalLimit(currentX, (fieldDimentionX - 1))) {
                continue;
            }

            // Check if current point is same at objetive.
            if (objectivePositionX == currentY && objectivePositionY == currentX) {
                // gameWindow.draw(getWall(currentX * sizeField, currentY * sizeField, sf::Color::Red));
                // gameWindow.display();
                // sf::sleep(sf::milliseconds(500));
                validMap = true;
                break;
            }

            // if (show) {
            //     gameWindow.draw(getWall(currentX * sizeField, currentY * sizeField, sf::Color::Yellow));
            //     gameWindow.display();
            //     sf::sleep(sf::milliseconds(20));
            // }

            if (!Game::arrayToMap[currentX][currentY - 1] && !Game::arrayToCheck[currentX][currentY - 1])
                    nodesQueue.push(pair<int, int>(currentX, currentY - 1));

            if (!Game::arrayToMap[currentX - 1][currentY] && !Game::arrayToCheck[currentX - 1][currentY])
                    nodesQueue.push(pair<int, int>(currentX - 1, currentY));

            if (!Game::arrayToMap[currentX][currentY + 1] && !Game::arrayToCheck[currentX][currentY + 1])
                    nodesQueue.push(pair<int, int>(currentX, currentY + 1));

            if (!Game::arrayToMap[currentX + 1][currentY] && !Game::arrayToCheck[currentX + 1][currentY])
                    nodesQueue.push(pair<int, int>(currentX + 1, currentY));

        }
        return validMap;
    }

    void Game::renderMapAndObjects() {

        gameWindow.clear();

        for (int i = 0; i < (fieldDimentionX - 1); i++) {
            for (int j = 0; j < (fieldDimentionY - 2); j++) {
                if (Game::arrayToMap[i][j]) {
                    gameWindow.draw(getWall(i * sizeField, j * sizeField, sf::Color::White));
                }
            }
        }

        playerObject.setPosition(playerPositionX * sizeField, playerPositionY * sizeField);
        objetiveObject.setPosition(objectivePositionX * sizeField, objectivePositionY * sizeField);

        gameWindow.draw(playerObject);
        gameWindow.draw(objetiveObject);
        gameWindow.display();
    }

    void Game::createNewMapAndObjects() {

        bool isValidMap = false;
        // int count = 0;
        while(!isValidMap) {

            Game::createNewMap();
            Game::setObjectPositions();
            
            // checkMapByBFS BFS, checkMapByDFS DFS.
            isValidMap = Game::checkMapByDFS(playerPositionY, playerPositionX, false);
            // for (int i = 0; i < (fieldDimentionX - 1); i++) {
            //     for (int j = 0; j < (fieldDimentionY - 2); j++) {
            //         Game::arrayToCheck[i][j] = false;
            //     }
            // }

            // Colorize on Yellow the steps.

            // if (isValidMap) {
            //     Game::renderMapAndObjects();
            //     sf::sleep(sf::milliseconds(20));
            //     isValidMap = Game::checkMapByBFS(playerPositionY, playerPositionX, true);
            //     count++;
            //     std::cout << "Cuenta: " + std::to_string(count) << std::endl;
            // }

            if (isValidMap) {
                isMapChecked = true;
            }
            // isValidMap = isMapChecked = false;
        }
    }

    void Game::run() {

        while(gameWindow.isOpen()) {

            if (!isMapGenerated) {
                createNewMapAndObjects();
                renderMapAndObjects();
                isMapGenerated = true;
            }
            processEvents();
        }
    }

int main() {

    Game game;
    game.run();
}
